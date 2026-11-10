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

### [x] A — Corriger les chemins de compilation — FAIT *(débloque choco)*

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

#### État de choco (10/11/2026)

`claire4 -cm` compile **les 15 fichiers de choco** (9 489 lignes), génère le
fichier système, et `go build` ne rapporte plus qu'**une** erreur.

Corrigés (commits `b0c1597`, `5cce29b`, `22459c4`, `cf7b45d`, `7b994db`, `8f52490`) :

| Erreur Go | Cause |
|---|---|
| `cannot use ...WriteFast(...) as Kernel.EID` | `WriteFast` renvoyait `*ClaireAny` |
| `undefined: Optimize` | `needed_modules` ne suivait que `uses` |
| `cannot use arg_N (EID) as *ClaireMethod` | `FASTCALLn` attendait une méthode |
| `...getActiveProblem_void().Id undefined` | méthode avec `RangeCheck` non marquée exception-ready |
| résultat faux (`GANY + 1` = 1 au lieu de 42) | lecture d'une globale `:any` produisait `EID{ptr,0}` |

**Reste** : `Choco/search.go:1173`
`cannot use try_6 (Kernel.EID) as int value in assignment`

### Analyse complète du dernier bug (11/11/2026)

**Symptôme.** Une méthode qui appelle une fonction pouvant lever est générée
comme fonction *native* ; le `g_try` qu'elle contient affecte alors un `EID`
à une variable `int` :

```go
func F_Repro_probe_void () int {          // native
  var Result int
  ...
  if ErrorIn(try_2) { Result = try_2 }    // EID dans int  ❌
```

**Cas minimal** : `test/regress/do-division`

```claire
[getF() : integer -> (AFIELD as Stored).latest]   // peut lever [15]
[probe() : integer -> getF() / 2]                 // générée native à tort
```

**Cause exacte** — `compile/osystem.cl` (et `src/Optimize/osystem.go`),
branche `Call_method` de `g_throw1` :

```claire
Call_method ( if (self.arg = m_unsafe | notOpt(self) = false | self.arg.selector = externC) false
              else if g_throw(self.args) true
              ... )
```

`notOpt(probe.body)` vaut `false` (division par un littéral non nul sûr),
donc la première branche répond **`false`** *sans jamais regarder les
arguments*. Or `getF()` peut lever : la division est sûre, pas son opérande.

**Pourquoi ce n'est pas corrigé tout de suite.** Tester les arguments en
premier (`g_throw(self.args)` avant le test `notOpt`) corrige
`do-division`, mais fait échouer choco à `var.cl:164` :

```
****[141] list<any>(dom, EID, "try_28", true, false) is a wrong arg list for Generate/g_statement
```

Le chemin `g_try` / `g_statement` (Itération + `when` + `case`, avec
`g_try_void`) ne supporte pas qu'une expression devienne « throwing »
là où elle ne l'était pas. Le correctif doit donc traiter les deux ensemble :

1. `g_throw(Call_method)` : regarder les arguments **avant** les cas
   d'optimisation ;
2. `g_statement` / `g_try` : supporter ce cas (actuellement
   `g_statement(dom, EID, "try_28", …)` reçoit une `Iteration` là où il
   attend une instruction feuille).

**Pièges rencontrés** (à ne pas refaire) :

- `notOpt(Call_method)` et `can_throw?` **mémorisent** le statut
  (`m.status := 0` optimiste dans `can_throw!`) : une analyse ratée fige le
  mauvais statut pour tous les appelants suivants ;
- `(x as C)` doit être considéré comme pouvant lever (`[15]`) quand
  `safety < 4` : `g_throw(Cast)` ne regardait que l'argument. Correctif
  isolé prêt (voir ci-dessous) mais **insuffisant seul** ;
- l'équilibrage des accolades de `do_statement`/`g_try` est fragile : j'ai
  cassé deux fois la génération en le touchant.

Les cas minimaux sont dans `test/regress/` (voir le README de ce dossier) :

| Cas | État |
|---|---|
| `gany-arithmetic` | ✅ corrigé (validé contre CLAIRE3 : 42/42/42) |
| `rangecheck-status` | ✅ corrigé (41 / 42) |
| `do-division` | ⬜ ouvert |

---

#### Historique : [ ] **F — `WriteFast` assigné à un `EID`**
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

**Version d'un module (v4.1.6)** : trois niveaux, par priorité —
`-publish <dir>@<version>`, puis le fichier `VERSION` à côté de `init.cl`,
puis `v0.0.0`. (CLAIRE4 n'a pas de slot `version` sur `module`.)

**Évolution proposée (à discuter avec Yves)** :
- [ ] ajouter un slot `version` à `module` (Kernel + `meta/object.cl` +
      `Core-meta.go`) : c'est la bonne abstraction, et cela permettrait à
      `use_module` de raisonner en compatibilité sémantique (`v1.x` → `v1.4.2`)
      au lieu de la comparaison lexicographique actuelle.

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

### [x] D — Go workspaces + `claire4 -install` — FAIT (commit `700ef6a`)

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

---
---

# Améliorations issues d'xlClaire (branche "améliorations" de CLAIRE3)

Trois chantiers issus de la branche d'améliorations d'xlClaire, à traiter après
le pipeline de modules.

---

## [ ] G — Options de ligne de commande déclarées par les modules

**Origine** : `xlclaire meta/cmdline.cl` (`option_respond`, `option_usage`,
`option_parsed`).

**Principe** : un module déclare comment répondre à ses propres options et
comment les documenter ; le parseur de la ligne de commande les prend en
compte automatiquement, et `-h` construit l'aide à partir de ces déclarations.

```claire
[option_respond(opt:{"--greet"}, l:list) : void ->
    let name := l[1] in
       (l :<< 1,                          // l'handler consomme ses arguments
        printf("Bonjour ~A\n", name))]

[option_usage(opt:{"--greet"}) : tuple(string, string, string) ->
    tuple("Dire bonjour",
          "--greet <nom:string>",
          "Affiche un message d'accueil.")]

[option_parsed() : void -> ... ]      // appelé une fois, après l'analyse
```

`option_respond` reçoit la liste des arguments restants où l'option a déjà été
retirée ; c'est l'handler qui retire ceux qu'il consomme (`l :<< n`).

**Implémentation** : le `while (l)` de `complex_main` (`compile/gomain.cl`,
`src/Generate/gomain.go`) teste les options built-in dans un `case`. Il suffit
d'insérer **avant** le `case` un appel aux restrictions de `option_respond`
(`ClaireProperty.Restrictions`), et après la boucle un appel à `option_parsed`.

**Décisions prises** :
- les options d'un module **s'ajoutent** aux options built-in ; en cas de
  collision, **le module gagne** (comme xlClaire), avec un trace(0) signalant
  l'écrasement ;
- `option_usage` construit `-h` : l'aide existante (`printHelp`) devient un jeu
  de restrictions `option_usage` dans Core/Reader/Generate, et `-h` les
  assemble en colonnes ;
- `option_parsed()` est appelé une seule fois, après la boucle, avant `%init?` ;
- l'API est exposée en `claire/` (pas dans un module dédié) pour que tout
  module puisse la redéfinir sans dépendance supplémentaire.

**Effort estimé** : 2-3 jours. **Risque** : faible (point d'extension, pas de
changement de sémantique).

**Validation** : un module `Repro` déclarant `--greet` compile, `claire4 -cm
Repro --greet Xavier` affiche le message, et `-h` liste l'option.

---

## [ ] H — Ports filtrables (`device` / `filter`)

**Origine** : `xlclaire meta/port.cl`, la partie « ports avec filtres ».

**État de CLAIRE4** : inexistant. `ClairePort`
(`src/Kernel/ClKernel.go:243`) est une struct **fermée** :

```go
type ClairePort struct {
    ClairePrimitive
    file   *os.File      // une seule sorte de source
    buffer []byte
    nChar, total, size, nEof int
    status int           // 0,1 fichier | 2,3 string | 4 stream
    firstc rune
}
```

et `Getc`/`Putc` dispatchent sur `p.status` dans un `switch`. Il n'y a ni
`device`, ni `filter`, ni `target`, ni `read_port` polymorphe.

**Ce qu'xlClaire avait** : deux niveaux, avec une interface à 6 méthodes
(`read_port`, `write_port`, `unget_port`, `eof_port?`, `flush_port`,
`close_port`). Un `filter` qui ne surcharge rien **relaie** vers
`self.target` ; un `device` répond lui-même ou signale l'absence
d'implémentation. D'où trois familles composables : sources (fichier, stdin,
socket), filtres (gzip, chiffrement, TLS), et le port manipulé par le code.

**Décision prise : ne pas réimplémenter `read_port`, utiliser `io.Reader` /
`io.Writer` de Go.** C'est le même principe (composition de flux) avec
l'écosystème standard au lieu d'une couche à maintenir.

```go
type ClairePort struct {
    ClairePrimitive
    reader io.Reader   // au lieu de *os.File
    writer io.Writer
    // le buffer existant reste, il lit via reader
}

// tous les filtres deviennent gratuits :
MakePort(bufio.NewReader(os.Stdin))
MakePort(readline.NewSource())                   // historique + complétion
MakePort(bufio.NewReader(gzip.NewReader(f, rc))) // decompression
MakePort(tlsConn)                                // TLS
```

**Étapes** :
- [ ] H1 — `ClairePort` passe à `io.Reader`/`io.Writer` (aucun changement
      visible ; validé par les tests existants et la compilation de choco).
- [ ] H2 — readline comme `io.Reader` (adaptateur au-dessus de `chzyer/readline`
      ou `peterh/liner`, **pur Go**, sans cgo), derrière un build tag :
      `go build -tags noreadline` doit continuer à produire un binaire sans
      dépendance externe.
- [ ] H3 — complétion sur les symboles du module courant (facultatif).

**Décisions prises** :
- **pas de cgo** : un binaire unique, sans dépendance système, est un acquis
  du projet (`go.work`, `claire4 -install`) ; lier `libedit` le casserait ;
- readline accessible par build tag pour que Yves puisse refuser la dépendance
  sans perdre le reste ;
- le mode non interactif (pipe, `< /dev/null`, CI) **ne doit jamais** activer
  readline : détection `term.IsTerminal(os.Stdin.Fd())` ;
- les ports mémoire (`status 2/3`) et fichier (`status 0/1`) gardent leur
  comportement actuel en H1 ; on n'introduit pas `device`/`filter` comme
  classes CLAIRE, la composition se fait en Go.

**Avant `io.Reader` seul, note** : `ClairePort` a aussi un `unget` implicite
(`firstc rune`, `F_pushback_port`). Il faut vérifier que le passage à
`io.Reader` conserve cette possibilité (un `bufio.Reader.UnreadRune` ou un
buffer interne le permet).

**Effort estimé** : H1 = 2-3 jours, H2 = 2-3 jours, H3 = +3 jours.
**Risque** : moyen pour H1 (cœur du Kernel I/O), faible ensuite.

**Validation** :
- H1 : `claire4`, `claire4 -cm choco`, les tests `test/regress` inchangés ;
- H2 : historique et complétion dans le REPL, et `go build -tags noreadline`
  produit toujours un binaire fonctionnel.

---

## [ ] I — Récapitulatif : ce qui manque à CLAIRE4 venu d'xlClaire

| Fonction xlClaire | État dans CLAIRE4 | Chantier |
|---|---|---|
| `use_module` + `CLAIRE_LIBS` | ✅ fait (B) | — |
| `-publish` / `-export` | ✅ `-publish` fait (C) ; `-export` non | C-bis |
| Options par module | ❌ | **G** |
| Ports filtrables | ❌ | **H** |
| readline / editline | ❌ | **H2** |
| Complétion | ❌ | **H3** |
| `save_context` / `restore_context` | ❌ (le REPL garde l'ancien modèle) | non planifié |
| Documentation auto (`codedoc`) | ❌ | non planifié |

---
---

# Nouveau cap : services REST de calcul (novembre 2026)

**Changement de stratégie** : l'UX et le middleware passent à Node/React.
CLAIRE4 ne fait plus que du **calcul**, exposé en REST, et sert de canal
typé entre process.

```
  Node / React  ──JSON──▶  dispatcher  ──exec + stdin──▶  claire4 (1 process / calcul)
                                  ◀──stdout──┘                    │
                                                                  └─ noyau, libs Go
```

**Décisions structurantes**

| Décision | Raison |
|---|---|
| **Pas de serveur HTTP dans CLAIRE** | le dispatcher (nginx/Node/Go) possède le port |
| **Un process par calcul** | le noyau n'est pas réentrant (pile d'évaluation globale) |
| **Modèle CGI** | `stdin` + variables d'environnement en entrée, `stdout` en sortie, comme `Http`/`Harbor` en xlClaire |
| **`ffor` / `fork` abandonnés** | `fork()` n'existe pas utilisable en Go ; déjà absent de CLAIRE4 |
| **JSON pour l'extérieur, `serialize` pour le canal CLAIRE↔CLAIRE** | le format binaire préserve classes et graphes partagés |

**Périmètre de la toolkit ClaireToolKit qui reste**

| Gardé | Abandonné |
|---|---|
| `Json`, `Regex`, `Md5`, `Sys` | `Wcl`, `WclSite`, `Pier`, `Harbor` (réécrits autrement) |
| l'esprit de `Http` (parsing CGI) et `Harbor` (`api(...)`) | `Zlib`, `Openssl` (Go les fournit) |
| *(à décider)* `Redis` + un client TCP | `Db`, `Dbo`, `Mysql`, `Xmlo`, `Sax`, `Dom`, `Dom3`, `Soap`, `Xmlrpc`, `Mail`, `Postal`, `WebLocale`, `LibLocale`, `Pdf`, `Gd`, `Gantt`, `Casper`, `Iconv` |
| `serialize`/`unserialize` (réécrit en Go) | la persistance compatible xlClaire (aucune donnée persistée) |

**Ce que `Pier` et `Harbor` apportaient et qu'on garde**

- `Pier` : passerelle JSON ↔ objets CLAIRE typés (`from_json(s, class)`).
  La partie `Dbo` (chargement d'objets depuis la base) disparaît avec la base.
- `Harbor` : lancer un serveur avec `-listen-port 8080` et déclarer les routes
  par des méthodes `api(...)`. Le dispatch par type set constant
  (`http_method:{"post"}, path:{"compute"}`) est conservé ; l'écoute du port
  passe au dispatcher.

---

## [ ] Performance : cibles à 10 Mo par requête

Les requêtes de calcul transportent des JSON de **plus de 10 Mo**. Mesures
faites sur le binaire actuel (`claire4 -cm`) :

| Test | Temps |
|---|---|
| démarrage + 10 Ko de sortie | 0,38 s |
| démarrage + 1 Mo de sortie | **1,5 s** (6,7 s CPU cumulé, GC actif) |
| projection 10 Mo | **~15 s** ❌ |

### Causes identifiées

**1. `Putc` sur stdout n'est pas bufferisé** — `src/Kernel/ClString.go:1135`

```go
} else if p.file == os.Stdout {
    p.file.Write([]byte(string(c)))   // un write() système par caractère
}
```

10 Mo de sortie = 10 millions d'appels système.

**2. `Getc` sur stdin décode une rune à la fois** — `ClString.go:1016`

```go
char, _, err := ToGoPort(p.Id()).reader.ReadRune()
```

Décodage UTF-8 complet + construction d'un `EID` par caractère.

**3. `BUFSIZE = 200`** — `ClString.go:889`

```go
var BUFSIZE int = 200   // test impact on speed (200 is the good value)
```

50 000 remplissages de buffer pour 10 Mo.

### Décision : passer du streaming au **batch**

```
AVANT (streaming)                 APRÈS (batch)
─────────────────                 ────────────
getc(stdin) × 10 M                read_all(stdin)   1 lecture
  ↓                                 ↓
parse par caractère               parse par index sur string
  ↓                                 ↓
Putc(cout) × 10 M                 construire la réponse, 1 écriture
```

`substring(s,i,j)` est une simple slice en Go : un parser qui avance par index
sur une string est 10 à 100× plus rapide qu'un parser à base de `getc`. Et un
`CLAIRE_CONTENT_LENGTH` connu (CGI) rend le streaming inutile.

**Gain attendu** : de ~15 s à **~0,3 s** pour 10 Mo.

### Travaux

- [ ] **R1** — bufferiser `Putc` y compris pour `stdout` ; flusher
      explicitement après le prompt du REPL et avant `exit`. *(0,5 j)*
- [ ] **R2** — `read_all(port)` : lecture par blocs (`io.ReadAll`), exposée en
      CLAIRE ; remplacer la boucle `getc` des parsers. *(0,5 j)*
- [ ] **R3** — les parsers JSON / HTTP avancent par index sur une string, jamais
      par `getc`. *(inclus dans J/T)*
- [ ] **R4** — benchmark de référence : générer un JSON de 10 Mo, mesurer temps
      et **pic mémoire** (un objet CLAIRE a un en-tête ; 10 Mo de JSON peuvent
      produire 100+ Mo d'objets). Borner avec `GOMEMLIMIT` si nécessaire. *(0,5 j)*

### Corrections annexes du Kernel (prérequises au mode process-par-calcul)

- [ ] `exit()` doit **flusher** `cout` avant `os.Exit`
      (`F_CL_exit`, `src/Kernel/ClEnv.go:451`) — aujourd'hui il ne flushe pas,
      donc la fin de la réponse peut être perdue.
- [ ] `exit()` ajoute un `"
"` parasite (`fmt.Printf("\n")`) : le conditionner
      au mode interactif, sinon il corrompt un corps JSON.

---

## Chantiers du nouveau périmètre

| # | Chantier | Contenu | Effort | Priorité |
|---|---|---|---|---|
| **R** | Perf + E/S | R1-R4 ci-dessus | 1,5 j | **1** |
| **J** | `Json` en Go | `encoding/json`, `from_json(s, class)` typé | 1,5 j | **1** |
| **T** | `Http` CGI | parsing requête (méthode, chemin, headers, query, body), écriture de la réponse ; port de l'esprit de `Http/source/{form,http_filter}.cl` | 2-3 j | **1** |
| **S** | Contrat stdin/stdout | enveloppe JSON + un exemple complet + doc | 1 j | **1** |
| **U** | `Harbor` | dispatch `api(req, method, path…)`, auth bearer | 1-2 j | 2 |
| **O** | `Regex`, `Md5`, `Sys` | `regexp`, `crypto/md5`, log/env | 1 j | 2 |
| **P** | `serialize` (canal CLAIRE) | format Go versionné, graphes partagés, cycles | 2-3 j | 2 |
| **M/N** | TCP + `Redis` | client socket puis RESP | 2-4 j | 3 |

**Priorité 1 = ~5-6 jours** pour un service de calcul opérationnel à 10 Mo.

### Questions ouvertes

- [ ] **Protocole CGI** : fastcgi (nginx), CGI classique, ou HTTP brut sur
      `stdin` (sans variables d'environnement) ? Le plus simple à tester est le
      HTTP brut ; le plus standard est CGI.
- [ ] **Enveloppe d'échange** : `{"method":…,"path":…,"body":…}` sur stdin et
      `{"status":…,"body":…}` sur stdout (permet au dispatcher de router), ou
      corps nu dans les deux sens ?
- [ ] **`Redis`** : nécessaire, ou le cache reste côté Node ?
- [ ] **`serialize`** : seulement CLAIRE → CLAIRE, ou aussi pour des données
      échangées avec Node (dans ce cas JSON suffit) ?
