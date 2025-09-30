## Using expr-lang in JavaScript environments (POC/Playground)

This repo is a practical POC trying out various ways to use expr-lang in JavaScript environments.

[Expr-lang](https://expr-lang.org) is a Go-centric expression language. You might have chosen it for your Go server and then discover you need to use it in the browser or another JavaScript environment (for example, to validate expression inputs).

* Root directory contains the example main Go package that validates expr-lang expressions against a specific [Env](https://expr-lang.org/docs/environment) schema.
* [wasm](./wasm) directory contains an example that compiles the Go package to WebAssembly using standard Go runtime and successfully uses it in both browser and Node.js. ✅
* [wasm-tinygo](./wasm-tinygo) directory contains an example that compiles the Go package to WebAssembly using [TinyGo](https://tinygo.org/) runtime and demonstrates that expr-lang is not compatible with TinyGo. ❌
  * This is confirmed by the maintainer of the expr-lang library here in [this GitHub issue](https://github.com/expr-lang/expr/discussions/451?utm_source=chatgpt.com#:~:text=(Note%20it%20is%20not%20possible%20to%20build%20reflect%20package%20with%20tinygo%20at%20the%20moment%2C%20as%20some%20reflect.Func%20bugs%20in%20tinygo).)


