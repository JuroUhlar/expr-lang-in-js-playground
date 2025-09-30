## Compiling your Expr-lang based Go package to WebAssembly

How recreate the setup in this folder: 

1. Create a [../main_js.go](../main_js.go) file in your go package directory. This should include `main` function that exposes your Go functions to JavaScript.

2. Run [./build_wasm.sh](./build_wasm.sh) to build the WebAssembly binary.
3. This will create a [main.wasm](./main.wasm) file in the `wasm` directory.
4. Run `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .` to copy the [wasm_exec.js](https://go.dev/wiki/WebAssembly?ref=blog.mozilla.ai#:~:text=To%20execute%20main.wasm%20in%20a%20browser) to the `wasm` directory. This is required to include in your JS code to run the WebAssembly module.
5. Write JavaScript code that imports the [wasm_exec.js](./wasm_exec.js) and [main.wasm](./main.wasm) files and runs the WebAssembly module.

* Example simple JS file: [test.js](./test.js).
* Example simple HTML file: [test.html](./test.html).
* Example simple JS server, serves the WASM binary compressed with gzip: [test-web-server.js](./test-web-server.js).
  * No dependencies here for simplicity.
  * A different compression method might be even better (Brottli?)
  * Check the network tab to see the that trasnferred size is 1,6 MB (vs 6.4 MB), 26% of original file size.

## Sources

* https://go.dev/wiki/WebAssembly?ref=blog.mozilla.ai
* https://www.sitepen.com/blog/compiling-go-to-webassembly


