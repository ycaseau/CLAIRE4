# SIGSEGV on `claire4 -cm choco` — `static_type(Call+)` returns a parameterized type

## Summary

Compiling any module that contains a `let` binding an `If` whose branches have
a **parameterized type** (e.g. `list<IntVar>`) makes the interpreter crash with
a nil-pointer dereference:

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x18]

Kernel.(*ClaireList).ValuesO(...)        ClKernel.go:742
Kernel.(*ClaireList).Length(...)         ClBag.go:613
Core.F_meet_class(...)                   Core/function.go:895
Language.F_static_type_any(...)          Language/pretty.go:1474
Language.(*Let).LetJito(...)             Language/define.go:3815
```

Found while porting **choco** (ClaireToolKit), at `choco/source/search.cl:414`.

## Minimal reproduction (7 lines, no choco needed)

```claire
C <: object(s:list<integer>, t:list<integer>)
[test(x:C) : void -> let l := (if x.s x.s else x.t) in print(length(l))]
```

```bash
claire4 -n -f repro.cl
```

→ panic.

Annotating the `let` avoids it, which confirms the `jito` path:

```claire
[test(x:C) : void -> let l:list<integer> := (if x.s x.s else x.t) in ...]   // OK
```

## Root cause

`static_type` is declared with range `class`, and branches `Call`, `Call_slot`
and `Call_method` all wrap their result with `class!`. The **`Call+` branch
does not**:

`meta/pretty.cl:444-445`

```claire
Call+ let p := self.selector, s := (p @ static_type(self.args[1])) in
         (case s (slot s.range, any p.range)),          // <- no class!
```

So `static_type(x.s)` where `s:list<integer>` returns the **Param**
`list<integer>` instead of the class `list`.

The chain is then:

1. `LetJito` (implicit typing of an unannotated `let`) calls
   `static_type` on the `If` — `Language/define.go:3815`;
2. `static_type(If)` computes
   `meet(static_type(arg), static_type(other))` — `Language/pretty.go:1474`;
3. `meet` takes two `class` arguments, casts them to `*ClaireClass` and reads
   `.Ancestors`;
4. a `Param` is not a `ClaireClass`, so `.Ancestors` is nil → **SIGSEGV**.

`meet(class!(x), class!(x))` on the same class returns `list` correctly, so the
problem is purely the missing `class!`.

## Fix

`meta/pretty.cl` (and the generated `src/Language/pretty.go`):

```diff
 Call+ let p := self.selector, s := (p @ static_type(self.args[1])) in
-         (case s (slot s.range, any p.range)),
+         (case s (slot class!(s.range), any class!(p.range))),
```

Two lines, no semantic change: `class!` on a class is the identity, and on a
parameterized type it gives its root class — which is what `meet` expects.

Available as commit `9a31455` on branch `fix/static-type-call-plus`
(fork `xpech/CLAIRE4`).

## What this unblocks

With the fix applied, `claire4 -cm choco` no longer crashes and the compiler
reports the real porting issues instead (e.g. `final(IntVar)` rejected as
"ephemeral classes cannot be final"), which are now tractable one by one.
