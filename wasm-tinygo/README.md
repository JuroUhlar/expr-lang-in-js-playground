# TinyGo WASM Attempt

## Summary

**Status**: ❌ **Does not work** - `expr-lang/expr` library incompatible with TinyGo

**Binary Size Comparison**:
- Standard Go (`wasm`): **5.9 MB**
- TinyGo (`wasm` target with `-scheduler=none`): **2.0 MB** (66% reduction)

## The Problem

The `expr-lang/expr` library does **not work** with TinyGo WASM due to reflection limitations.

### Error Output
```
Received expression: bot == true len: 11
Error compiling: unexpected token Operator("==") (1:5)
 | bot == true
 | ....^
```

### Root Cause

1. **TinyGo's Limited Reflection Support**:
   - TinyGo doesn't fully support Go's reflection package
   - The `expr-lang/expr` library relies heavily on reflection for parsing expressions
   - The parser fails to recognize basic operators like `==`, `>=`, etc.

2. **WASM Integration** (Solved):
   - ✅ Successfully integrated TinyGo's `//export` mechanism
   - ✅ Properly pass strings via pointer + length from JavaScript
   - ✅ WASI initialization works correctly
   - ❌ But the expr library itself doesn't function

## What Was Accomplished

1. ✅ Created TinyGo-specific build with `//export` directives
2. ✅ Implemented proper string marshalling (JS → WASM memory → Go string)
3. ✅ Integrated WASI support with Node.js
4. ✅ Binary compiles and exports are callable from JavaScript
5. ❌ The expr parser doesn't work due to reflection limitations

## Technical Details

### Build Configuration
- Target: `wasm` with `-scheduler=none`
- Export mechanism: `//export` directives (not `syscall/js`)
- String passing: Manual memory allocation via `malloc`/`free`

### Files Created
- `main_tinygo.go` - TinyGo-specific implementation with `//export`
- `build_wasm.sh` - Build script using TinyGo compiler
- `test.js` - Node.js test with proper WASI + string marshalling
- `test.html` - Browser test (not tested, same underlying issue)
- `wasm_exec.js` - TinyGo's WASM runtime

## Conclusion

While TinyGo can produce smaller WASM binaries (2.0 MB vs 5.9 MB), it **cannot** be used with the `expr-lang/expr` library due to fundamental reflection incompatibilities. The library's parser requires reflection features that TinyGo doesn't fully support.

**Recommendation**: Use standard Go WASM for this use case. The expr library works perfectly with standard Go.
