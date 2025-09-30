# Compiling your Expr-lang based Go package to WebAssembly

How recreate the setup in this folder: 

1. Create a [../main_js.go](../main_js.go) file in your go package directory. This should include `main` function that exposes your Go functions to JavaScript.

2. Run [./build_wasm.sh](./build_wasm.sh) to build the WebAssembly binary.
3. This will create a [main.wasm](./main.wasm) file in the `wasm` directory.
4. Run `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .` to copy the [wasm_exec.js](https://go.dev/wiki/WebAssembly?ref=blog.mozilla.ai#:~:text=To%20execute%20main.wasm%20in%20a%20browser) to the `wasm` directory. This is required to include in your JS code to run the WebAssembly module.
5. Write a JS file that imports the [wasm_exec.js](./wasm_exec.js) and [main.wasm](./main.wasm) files and runs the WebAssembly module: [test.js](./test.js).
6. Run the JS file: `node test.js`.

## Sources

* https://go.dev/wiki/WebAssembly?ref=blog.mozilla.ai
* https://www.sitepen.com/blog/compiling-go-to-webassembly


