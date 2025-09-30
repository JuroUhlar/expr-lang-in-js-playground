const fs = require('fs');
require('./wasm_exec');

const go = new Go();
const wasmBuffer = fs.readFileSync(__dirname + '/main.wasm');

const exampleExpressions = [
    'bot == true',
    'suspectScore >= 50',
    'bot == unknownField',
    'bot != 12',
    'unknownField == true',
];

WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    go.run(result.instance);
    
    for (const expr of exampleExpressions) {
        console.log('Expression:', expr);
        console.log('Is valid:', isExpressionValid(expr));
        console.log();
    }
});
