# CLAIRE4 — Feuille de route build & modules

> Document de travail (branche `fix/choco-crash`).
> Objectif : rendre CLAIRE4 utilisable comme CLAIRE3/xlClaire en production —
> installation simple, compilation de modules fiable, dépendances versionnées.

---

## Contexte

Deux problèmes constatés en usage réel :

1. **Crash à la compilation du module `choco`** (`claire4 -cm choco` → SIGSEGV).
2. **Déploiement / compilation des modules trop rigide**, sans équivalent des
   mécanismes xlClaire (`configure && make`, `-publish`, `use_module`).

---

## Sujet 1 — Crash choco : RÉSOLU ✅

**Branche** : `fix/choco-crash` — **commit** `8dd9f10`

### Cause racine

Bug de `static_type` dans l'interpréteur (indépendant de choco).

| # | Étape | Emplacement |
|---|-------|-------------|
| 1 | `let lvar := (if a.varsToStore a.varsToStore else a.problem.vars)` — deux branches typées `list<IntVar>` | `choco/source/search.cl:414` |
| 2 | `LetJito` (typage implicite d'un `let` **non annoté**) appelle `static_type` sur le `If` | `src/Language/define.go:3815` |
| 3 | `static_type(If)` → `meet(static_type(arg), static_type(other))` | `src/Language/pretty.go:1474` |
| 4 | La branche **`Call+`** renvoyait `s.range` / `p.range` **sans `class!`** → un `Param` (`list<IntVar>`) au lieu d'une classe | `src/Language/pretty.go:1386-1390` |
| 5 | `meet` caste en `*ClaireClass`, lit `.Ancestors` → **nil → SIGSEGV** | `src/Core/function.go:895` |

Les branches `Call`, `Call_slot`, `Call_method` faisaient déjà `class!` ; seul `Call+` l'oubliait.

### Reproduction minimale (7 lignes)

```claire
C <: object(s:list<integer>, t:list<integer>)
[test(x:C) : void -> let l := (if x.s x.s else x.t) in print(length(l))]
```

→ même panic. Avec `let l:list<integer> := ...` (annoté) : pas de crash (confirme `jito`).

### Correctif

`meta/pretty.cl` (source) + `src/Language/pretty.go` (Go généré) :

```claire
Call+ let p := self.selector, s := (p @ static_type(self.args[1])) in
         (case s (slot class!(s.range), any class!(p.range))),   // <- class! ajouté
```

### Vérification

- Repro minimal : OK (plus de panic).
- `claire4 -cm choco` : plus de SIGSEGV, la compilation démarre réellement
  (`==== Generate Go files for module choco` / `++++ Compiling the file chocutils.cl`).

### Reste à faire côté choco (pour Yves / portage)

- `final(IntVar)`, `final(LinComb)`, `final(Occurrence)` échouent en CLAIRE4
  (`compil.cl:74,123,129`) : la classe a `open = 2` (héritage `Ephemeral`).
  CLAIRE3 tolérait `final` sur `open=2`. Ici `EXTENDABLE_CHOCO = false`,
  les trois `final(...)` peuvent être retirés.
- Ensuite on bute sur le chemin codé en dur → **Sujet 2, point A**.

---

## Sujet 2 — Build, installation, dépendances

### Diagnostic du modèle actuel

Tout le build tient en 3 lignes de shell dans `compile/gomain.cl:282-292` :

```claire
[compile_dir(m:module): void
 -> let s := "mkdir -p " /+ home() / "go" / "src" / capitalize(string!(m.name)) in shell(s)]

[compile_exe(%out:string): void
 -> let s := "go build " /+ outdir /+ home() / "go" / "src" / %out /+ ".go" in shell(s)]
```

| Problème | Conséquence |
|----------|-------------|
| `home()` = `$CLAIRE_HOME`, puis concaténation `/go/src` **en dur** | chemin faux si `CLAIRE_HOME` vaut déjà `.../go/src` |
| `compiler.source` figé à la compilation du binaire | `cannot open file /Users/ycaseau/claire/v4.1/go/src/Choco/chocutils.go` |
| GOPATH global obligatoire (`GOPATH` + `GO111MODULE=off`) | pas de projet isolé, pas de CI simple |
| Pas de `go.mod` par module | écosystème Go inaccessible |
| Tous les modules dans `$CLAIRE_HOME/../go/src` | un module ne peut pas vivre dans son propre dépôt |
| Pas de `use_module` | édition manuelle de `init.cl` |
| Pas de `-publish` | aucun module consommable comme dépendance |
| Installer = copier des répertoires | installation fragile |

### Référence : ce que fait xlClaire

| Mécanisme | Où (xlClaire) | Rôle |
|---|---|---|
| `configure && make` | `Makefile.unix.in`, `install.cl` | build complet / incrémental |
| `-publish` / `-export` | `csrc/Generate.cpp:7822+` | copie `init.cl`, `include/`, `source/`, `.o/.so` dans `$CLAIRE_HOME/lib/<m>/<version>-<release>-<env>/` |
| `use_module("m/version")` | `meta/file.cl:543` | cherche dans `CLAIRE_LIBS` + `$CLAIRE_HOME/lib`, charge `init.cl` puis le module (compilé si possible, sinon interprété) |

---

## Actions (ordonnées)

### [~] A — Corriger les chemins de compilation *(priorité : débloque choco)*

**Décision retenue : modèle "style Cargo" via `go.work`** (testé et validé).

#### Pourquoi cette option

| | Workspace Go (Cargo-like) | GOPATH local `./build/` | Préfixe `claire4/Kernel` |
|---|---|---|---|
| Imports générés | **inchangés** | inchangés | à réécrire (29 fichiers + générateur) |
| ISO écosystème Go 2020+ | oui | non (déprécié) | oui |
| Dépendances externes Go | oui | non | oui |
| Verrouillage versions | oui (`go.sum`) | non | oui |
| Risque régression | **faible** | faible | élevé |

Test de validation effectué (`/tmp/gowork`) : trois modules (`Kernel`, `Core`, `Choco`)
reliés par `go.work` + `replace`, imports **sans préfixe** (`. "Kernel"`, `"Core"`),
`go build` OK depuis le sous-répertoire du module.

#### Correspondance Cargo ↔ Go

| Cargo | Go |
|---|---|
| `Cargo.toml` (par crate) | `go.mod` (par module) |
| `Cargo.lock` | `go.sum` + `go.work.sum` |
| `[workspace] members` | `go.work` + `use (...)` |
| `path = "../core"` | `replace Kernel => ../core` |
| `target/` | `$GOCACHE` |
| `cargo build -p choco` | `go build ./choco` |
| *(pas d'équivalent)* | `go.sum` verrouillage transitif |

#### Cible

```
<module>/
├── go.mod                # module Choco (écrit/mis à jour par claire4)
├── init.cl
├── source/*.cl
└── build/
    ├── go.work           # use (./Choco + modules de base)
    └── Choco/*.go        # PRODUCER.source
```

#### Travaux

- [x] A1 — `compile_dir` crée `compiler.source` **et** `compiler.source/<Module>` — commit `8ef32ff`
- [x] A2 — `compile_exe` : `cd <compiler.source> && go build -o <out> ./<Module>` — commit `8ef32ff`
- [x] A3 — `prepare_workspace()` écrit `go.work` + un `go.mod` par module de base — commit `8ef32ff`
- [x] A4 — `compiler.source` = `$CLAIRE_SOURCE` > `$CLAIRE_HOME` > `.` — commit `8ef32ff`
- [x] A5 — erreur `[213]` si `compiler.source` est vide — commit `8ef32ff`

**Résultat sur choco** : les 15 fichiers sont générés dans
`$CLAIRE_SOURCE/Choco`, `go.work` est écrit, `go build` résout tous les
paquets de base (`Kernel`, `Core`, `Language`, `Reader`, `Optimize`,
`Generate`).

#### Reste à traiter (nouveau, hors A)

- [ ] **F — `WriteFast` assigné à un `EID`** : `go build` de choco échoue sur
  des `X.WriteFast(...)` là où du code `EID` est attendu
  (`Choco/chocapi.go:357`, `prop.go:2481`, …).
  `ClaireProperty.WriteFast` renvoie `*ClaireAny` alors que le contexte attend
  un `EID`. À trancher : soit le générateur doit émettre `WriteEID`
  (comme pour `write_fast`, cf. `compile/gostat.cl:834`), soit
  `WriteFast` doit renvoyer un `EID`.
- [ ] **E — EOF** : `claire4 … </dev/null` boucle indéfiniment dans le REPL
  (gêne les tests automatisés ; masquait aussi l'erreur `WriteFast`).

- **Fichiers** : `compile/gomain.cl`, `src/Generate/gomain.go`, `src/Optimize/Optimize-meta.go`
- **Risque** : faible (imports générés inchangés)
- **Validation** : `claire4 -cm choco` produit `build/Choco/chocutils.go`, puis erreur suivante
  (attendu : `final(...)` côté choco, puis d'éventuelles erreurs de portage).

#### Limite connue (à traiter en C)

`replace` et `go.work` référencent des **chemins locaux**. Pour qu'un module publié
soit consommable ailleurs, il faudra soit publier les modules de base, soit figer
par `go mod vendor` — c'est le rôle du `-publish` (point C).

### [x] B — `use_module("m/version")` — FAIT

**Commits** : `864dd6e` (primitives) + `7aae723` (use_module)

- **Fichiers** : `meta/file.cl` ↔ `src/Reader/file.go` + `src/Reader/Reader-meta.go`,
  `meta/function.cl` ↔ `src/Core/Core-meta.go`, `src/Kernel/ClString.go` (primitives)
- **Comportement implémenté** :
  1. résolution : `$CLAIRE_LIBS` puis `$CLAIRE_HOME/lib` puis répertoire courant
  2. sans version explicite : plus grande version disponible (ordre lexicographique)
  3. `chdir` dans le répertoire du module, `load("init")`, puis `load(m)`
  4. enregistrement dans `used_modules` (liste de paires module → répertoire)
  5. warning listant tous les répertoires fouillés si le module est absent
- **Primitives ajoutées** (Kernel) : `isfile?`, `isdir?`, `entries`, `pwd`, `setcwd`
- **Validation** : `CLAIRE_LIBS=./libs` + `use_module("mymod")` résout `v2.1.0`,
  `mymod/hello()` s'exécute, `print_used_modules()` affiche la résolution.

**Reste à faire pour B+** :
- [ ] charger `build/` (code Go compilé) quand il existe, au lieu des sources
- [ ] écrire un `claire.lock` (ou l'équivalent `go.mod`) à partir de `used_modules`
- [ ] utiliser `used_modules` dans `-publish` (point C)

### [x] C — `-publish` — FAIT

**Commit** : `dda877b`

    claire4 -cm <module> -publish <dir>[@<version>] [-force]

Produit `$CLAIRE_LIBS/<name>/<version>/{init.cl, source/, build/<Module>/, META.json}`,
immédiatement consommable par `use_module`.

**Vérifié** : `-cm mymod -publish ./libs@v1.0.0` puis
`CLAIRE_LIBS=./libs claire4 -f use.cl` → `Hello from mymod!`.

**Non fait** :
- [ ] `-export` (publication dans un répertoire `CLAIRE_LIBS` choisi par index)
- [ ] figer les dépendances (`go mod vendor` / `claire.lock`) pour un module publié sans ses base modules

### [ ] C-bis — `-publish` / `-export` (notes d'origine)

- **Fichiers** : `compile/gomain.cl`, `src/Generate/gomain.go`
- **Quoi** : produire un dossier autosuffisant et versionné
  ```
  $CLAIRE_LIBS/<m>/<version>-<os>-<arch>/
      init.cl
      source/
      build/        # code Go généré
      go.mod        # dépendances figées
      META.json     # version, plateforme, hash, date
  ```
- **Risque** : moyen
- **Validation** : publier choco, puis le consommer depuis un autre répertoire via `use_module`.

### [ ] D — `go.mod` racine + `claire4 -install`

- **Fichiers** : nouveau `src/go.mod`, `compile/gomain.cl`
- **Quoi** :
  - `go.mod` racine dans `src/` → `go build ./...` sans `GOPATH` ni `GO111MODULE=off`
  - `claire4 -install` crée `~/.claire4/{lib,bin,init.cl}`
  - Core / Language / Reader / Optimize / Generate restent dans le binaire (rien à installer)
- **Risque** : faible
- **Validation** : `go build ./...` propre ; `claire4 -install` puis `claire4` démarre.

### [x] E — EOF : sortir proprement du REPL — FAIT (commit `cf6370f`)

- **Fichiers** : `meta/inspect.cl`, `src/Reader/inspect.go`
- **Quoi** : quand `nextunit(reader)` renvoie `Reader/eof`, sortir de la boucle
  (CLAIRE3 faisait `break()` ; CLAIRE4 boucle indéfiniment et noie les logs)
- **Risque** : faible
- **Validation** : `claire4 -cm choco </dev/null` se termine au lieu de tourner sans fin.

---

## Ordre de travail convenu

1. **A** (débloque choco, valide la chaîne complète), puis
2. **B** (le plus structurant : vrai système de dépendances), puis
3. **C**, **D**, **E**.
