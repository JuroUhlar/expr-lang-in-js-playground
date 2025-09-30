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
}).then((result) => {
    const wasm = result.instance;
    
    // Set up the Go runtime instance
    go._inst = wasm;
    
    // Call wasi.start() to initialize
    try {
        wasi.start(wasm);
    } catch(e) {
        // _start might exit immediately, that's ok for exports
    }
    
    console.log('TinyGo WASM Expression Validator Test\n');
    
    // Helper to allocate string in WASM memory
    function allocateString(str) {
        const encoder = new TextEncoder();
        const bytes = encoder.encode(str);
        const ptr = wasm.exports.malloc(bytes.length);
        const mem = new Uint8Array(wasm.exports.memory.buffer);
        mem.set(bytes, ptr);
        return { ptr, len: bytes.length };
    }
    
    exampleExpressions.forEach(expr => {
        // Allocate string in WASM memory
        const { ptr, len } = allocateString(expr);
        
        // Call the exported isValid function with pointer and length
        const isValid = wasm.exports.isValid(ptr, len);
        
        // Free the allocated memory
        wasm.exports.free(ptr);
        
        console.log(`${expr.padEnd(30)} → is valid: ${isValid === 1}`);
    });
}).catch(err => {
    console.error('Error:', err.message);
    console.error(err.stack);
    process.exit(1);
});
