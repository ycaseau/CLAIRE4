# CLAIRE4 — Crash à la compilation de choco

> Correctif du crash `static_type(Call+)` (SIGSEGV pendant `claire4 -cm choco`).
> Le reste du travail (build, modules, `-publish`, `use_module`) est sur une
> autre branche : `feat/modules-and-build`.


## Cause racine

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

---

## Vérification

```bash
cd ClaireToolKit/choco && claire4 -cm choco
```

Avant : `panic: runtime error: invalid memory address or nil pointer dereference`
dans `Core.F_meet_class`.

Après : plus de crash ; la compilation démarre et rapporte les vrais problèmes
de portage (`final(IntVar)` refusé, chemins de build), traités sur l'autre
branche.
