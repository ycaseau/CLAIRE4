# `claire4 -cm` crashes with a SIGSEGV: `static_type(Call+)` returns a parameterized type

## Summary

Compiling a module that binds an `If` to an unannotated `let`, when the two
branches have a **parameterized type** (e.g. `list<IntVar>`), makes the
interpreter crash with a nil-pointer dereference.

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x18]

Kernel.(*ClaireList).ValuesO(...)        ClKernel.go:742
Kernel.(*ClaireList).Length(...)         ClBag.go:613
Core.F_meet_class(...)                   Core/function.go:895
Language.F_static_type_any(...)          Language/pretty.go:1474
Language.(*Let).LetJito(...)             Language/define.go:3815
Language.(*Defmethod).SelfEval(...)      Language/define.go:967
Reader.F_load_file_string(...)           Reader/file.go:459
```

Found while porting **choco**, at `choco/source/search.cl:414`.

## Reproduction

```claire
C <: object(s:list<integer>, t:list<integer>)
[test(x:C) : void -> let l := (if x.s x.s else x.t) in print(length(l))]
```

```bash
claire4 -n -f repro.cl
```

Panics. With an explicit type on the `let`, it does not — which points at
the `jito` (implicit typing) path:

```claire
[test(x:C) : void -> let l:list<integer> := (if x.s x.s else x.t) in ...]   # OK
```

## Cause

`static_type` has range `class`, and the branches `Call`, `Call_slot` and
`Call_method` wrap their result with `class!`. The **`Call+` branch does not**
(`meta/pretty.cl:444`):

```claire
Call+ let p := self.selector, s := (p @ static_type(self.args[1])) in
         (case s (slot s.range, any p.range)),          // <- no class!
```

So `static_type(x.s)` with `s:list<integer>` returns the `Param`
`list<integer>` instead of the class `list`. Then:

1. `LetJito` calls `static_type` on the `If` (`Language/define.go:3815`);
2. `static_type(If)` computes `meet(static_type(arg), static_type(other))`
   (`Language/pretty.go:1474`);
3. `meet` takes two `class`, casts them to `*ClaireClass` and reads
   `.Ancestors` (`Core/function.go:895`);
4. a `Param` is not a `ClaireClass` → `.Ancestors` is nil → **SIGSEGV**.

## Fix

```diff
 Call+ let p := self.selector, s := (p @ static_type(self.args[1])) in
-         (case s (slot s.range, any p.range)),
+         (case s (slot class!(s.range), any class!(p.range))),
```

Two lines in `meta/pretty.cl` and the generated `src/Language/pretty.go`.
`class!` is the identity on a class and gives the root class on a
parameterized type, which is what `meet` expects.

Branch: `fix/static-type-call-plus` on `xpech/CLAIRE4`
(commit `9a31455`).

## Effect

With the fix, `claire4 -cm choco` no longer crashes and reports the real
porting issues instead (e.g. `final(IntVar)` rejected as "ephemeral classes
cannot be final"), which are tractable one by one.
