# TinyGo WASM Attempt

## Summary

**Status**: ❌ **Does not work** - `expr-lang/expr` library incompatible with TinyGo

**Binary Size Comparison**:
- Standard Go (`wasm`): **5.9 MB**
- TinyGo (`wasm` target): **2.5 MB** (58% reduction)

## The Problem

The `expr-lang/expr` library does **not work** with TinyGo WASM due to missing reflection methods.

### Error Output

Run the test to reproduce the error:
```bash
node test.js
```

Error output:
```
panic: unimplemented: (reflect.Type).MethodByName()
RuntimeError: unreachable
    at main.(*reflect.rawType).MethodByName
    at main.(*github.com/expr-lang/expr/checker.checker).ident
```

### Root Cause

**TinyGo's Incomplete Reflection Support**:
- TinyGo doesn't implement `reflect.Type.MethodByName()` 
- The `expr-lang/expr` library requires this method for type checking
- This seems to be a fundamental limitation of TinyGo's reflection implementation.

## What Was Accomplished

1. ✅ Successfully used `syscall/js` with TinyGo (simpler than `//export`)
2. ✅ Proper WASI integration with Node.js
3. ✅ Function exports work perfectly
4. ✅ String passing works (no manual memory allocation needed with `syscall/js`)
5. ❌ The expr library crashes on `MethodByName()` call


### Files
- `main_syscalljs.go` - Clean implementation using `syscall/js`
- `main_exports.go` - Alternative implementation using `//export` (more complex)
- `build_wasm.sh` - TinyGo build script
- `test.js` - Node.js test (shows the MethodByName error)
- `wasm_exec.js` - TinyGo's WASM runtime

## Conclusion

**Recommendation**: Use standard Go WASM for this use case. The expr library requires full reflection support.
