# Branch `feat/modules-and-build` — what it changes

> Fork `xpech/CLAIRE4`, based on `ycaseau/CLAIRE4@main` (`25b1496`).
> 42 commits. Two large areas: **the module build pipeline** and **code
> generation fixes found while porting choco**.

Each commit patches the generated Go file **and** its `.cl` source, so that
recompiling the compiler (`claire4 -cc Generate`) does not reintroduce the bug.
Files under `src/Kernel/` are handwritten Go and are patched there only.

---

## 1. The crash (also proposed as a standalone PR)

`fix(Language): static_type(Call+) must return a class, not a parameterized type`

See `BUG-static-type-call-plus.md`. Two lines in `meta/pretty.cl` +
`src/Language/pretty.go`. Without it, `claire4 -cm choco` dies with a SIGSEGV.

---

## 2. Build pipeline (actions A–E)

### A — Module output paths

**Before**: three inconsistent notions of "where the code goes".

```
compile_dir   : $CLAIRE_HOME + "/go/src/" + <Module>
compile_exe   : $CLAIRE_HOME + "/go/src/" + <out> + ".go"
generator     : compiler.source + "/" + <Module>
```

With the documented setting `CLAIRE_HOME=<root>/go/src`, the first two produced
`<root>/go/src/go/src/Choco`. With `CLAIRE_HOME` unset, `compiler.source` fell
back on a path hard-coded at build time (`/Users/ycaseau/...`).

**Now**:

- `compile_dir` creates `compiler.source` and `compiler.source/<Module>`;
- `compile_exe` builds `./<Module>.go` from `compiler.source` and writes the
  binary to `<module>/bin/<module>`;
- `compiler.source` comes from the environment: `$CLAIRE_SOURCE` >
  `$CLAIRE_HOME` > `.` (was hard-coded);
- a module without `source` is now a clear `[212]` error instead of an
  out-of-memory loop.

### B — `use_module("name/version")`

Transposed from xlClaire `meta/file.cl`. A module is published as a
self-contained directory:

```
$CLAIRE_LIBS/<name>/<version>/
    init.cl         -- module declaration (the equivalent of a .h)
    source/         -- CLAIRE sources, loaded by the interpreter
    build/<Module>/ -- generated Go package (the equivalent of a .so)
    META.json
```

Lookup order: `$CLAIRE_LIBS`, `$CLAIRE_HOME/lib`, `~/.claire4/lib`, current
directory. Without a version, the greatest version directory wins.

Also added: `module_libs()`, `find_module_dir()`, `print_used_modules()`,
`used_modules`, and the filesystem primitives they need (`isfile?`, `isdir?`,
`entries`, `pwd`, `setcwd`) in the Kernel.

### C — `-publish`

```
claire4 -cm <module> -publish <dir>[@<version>] [-force]
```

Copies the module into `<dir>/<name>/<version>/` with `init.cl`, `source/`,
`build/<Module>/` and `META.json`. Without a directory it falls back on
`$CLAIRE_LIBS`, then `$CLAIRE_HOME/lib`, then `~/.claire4/lib`. The version
comes from `@<version>`, then a `VERSION` file next to `init.cl`, then
`v0.0.0`. Refuses to overwrite without `-force`.

CLAIRE4 has no `version` slot on `module` (a free slot is rejected with
`[106]`), hence the file/option convention.

When compiling a module that uses another one, `prepare_workspace()` adds the
dependency's `build/<Module>` to `go.work`, so the compiled package is linked
as well as the sources being loaded by the interpreter.

### D — Go workspaces, and `claire4 -install`

The generated Go code imports the base packages by their bare name
(`. "Kernel"`, `"Core"`, …), which used to require `GOPATH` +
`GO111MODULE=off` and a copy of the tree under `$GOPATH/src`.

`src/go.work` + one `go.mod` per package make those short imports resolve
directly:

```bash
cd src && go build -o claire4 claire4.go
cd src && go build ./...
```

`claire4 -install` creates `$CLAIRE_HOME` (or `~/.claire4`) with `lib/`,
`share/claire4/` (the base packages) and `init.cl`, and copies the binary into
a directory of `$PATH` — `$CLAIRE_BIN`, `$GOBIN`, `~/.local/bin`, then
`<home>/bin`; `/usr/local/bin` only with `CLAIRE_BIN_SYSTEM=1`.

### E — End of input in the REPL

`nextunit` answers the symbol `Reader/eof` when the input is exhausted, but
`top_level` only stopped on `q`. The eof symbol was fed to `eval()` in a tight
loop, so `claire4 … </dev/null` never terminated and produced millions of
`Reader/eof` lines, hiding the real error underneath.

---

## 3. Code generation fixes found on choco

Choco compiles in 9 489 lines across 15 files. Getting there required five
fixes; each one is reproducible with a few lines (see `test/regress/`).

| Commit | Bug | Symptom |
|---|---|---|
| `0b4ceeb` | `WriteFast` returned `*ClaireAny` while every caller stores the result in an `EID` | `cannot use …WriteFast(…) as Kernel.EID` |
| `36f214a` + `274598f` | `FASTCALLn` took a `*ClaireMethod`, but the compiler emits the method as an `EID` | `cannot use arg_3 (EID) as *ClaireMethod` |
| `9418e3d` | `needed_modules` followed only `uses`, but second-order types compile to `Optimize/F_nth_property` | `undefined: Optimize` |
| `d1dbb6a` | a method ending with `RangeCheck` was not marked exception-ready | `F_getActiveProblem_void().Id undefined` |
| `29c1da3` | reading a global declared `:any` emitted `EID{Value,0}` (pointer in PTR, 0 in VAL) | `GANY + 1` returned **1** instead of 42 |

The last one is a **silent wrong-result** bug. Comparing with CLAIRE3
(`cclaire` from xlclaire-v3.7.8) is what proved it: CLAIRE3 returns 42.

### Still open

`Choco/search.go:1173` — `cannot use try_6 (Kernel.EID) as int value`.

Minimal case in `test/regress/do-division`:

```claire
[getF() : integer -> (AFIELD as Stored).latest]   // may raise [15]
[probe() : integer -> getF() / 2]                 // generated native, should be EID
```

Root cause, in `compile/osystem.cl` (`g_throw1`, branch `Call_method`):

```claire
Call_method ( if (self.arg = m_unsafe | notOpt(self) = false | …) false
              else if g_throw(self.args) true
              … )
```

`notOpt(/ by a non-zero literal)` is false, so the first branch answers "no
throw" **without ever looking at the operands**. The division is safe; its
operand is not.

Testing the arguments first repairs this case but breaks choco at
`var.cl:164` (`[141] list<any>(dom, EID, "try_28", …) is a wrong arg list for
Generate/g_statement`): `g_statement`/`g_try` do not yet support an expression
that becomes throwing where it was not. Both sides must be fixed together.

---

## 4. Regression cases

`test/regress/` — three self-contained modules, each documenting its bug in
`init.cl`:

| Case | State |
|---|---|
| `gany-arithmetic` | ✅ fixed (`29c1da3`), validated against CLAIRE3 |
| `rangecheck-status` | ✅ fixed (`d1dbb6a`) |
| `do-division` | ⬜ open — it is the purpose of the test |

```bash
cd test/regress/gany-arithmetic
claire4 -cm GanyArithmetic
./bin/ganyarithmetic
```

---

## 5. Not included

The branch also documents (but does not implement) the next steps, in
`TODO-CLAIRE4.md`:

- module command line options (`option_respond` / `option_usage`, from
  xlClaire);
- filterable ports (`device` / `filter`), replaced by a move of `ClairePort`
  to `io.Reader` / `io.Writer`;
- performance work for 10 MB requests: `Putc` is unbuffered for stdout
  (one `write()` syscall per character), `Getc` decodes one rune at a time,
  `BUFSIZE` is 200.
