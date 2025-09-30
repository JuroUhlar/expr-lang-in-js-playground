const fs = require('fs');
const { WASI } = require('wasi');
require('./wasm_exec');

const exampleExpressions = [
    'bot == true',
    'suspectScore >= 50',
    'bot == unknownField',
    'bot != 12',
    'unknownField == true',
];

const wasi = new WASI({
    version: 'preview1',
    args: process.argv,
    env: process.env,
});

const go = new Go();
const wasmBuffer = fs.readFileSync(__dirname + '/main.wasm');

WebAssembly.instantiate(wasmBuffer, {
  ...go.importObject,
  ...wasi.getImportObject(),
})
  .then(async (result) => {
    // Set up the instance for Go
    go._inst = result.instance;

    // Initialize WASI
    try {
      wasi.initialize(result.instance);
    } catch (e) {
      // May fail if _start exists, that's ok
    }

    // Start the Go runtime (non-blocking)
    try {
      go.run(result.instance);
    } catch (e) {
      // _start might exit, that's ok for our use case
    }

    // Wait a bit for initialization
    await new Promise((resolve) => setTimeout(resolve, 100));

    console.log("TinyGo WASM Expression Validator Test (syscall/js)\n");
    console.log("isExpressionValid available?", typeof isExpressionValid);

    exampleExpressions.forEach((expr) => {
      const isValid = isExpressionValid(expr);
      console.log(`${expr.padEnd(30)} → is valid: ${isValid}`);
    });
  })
  .catch((err) => {
    console.error("Error:", err.message);
    console.error(err.stack);
    process.exit(1);
  });
