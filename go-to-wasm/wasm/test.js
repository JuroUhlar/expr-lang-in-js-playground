const fs = require('fs');
const path = require('path');
require('./wasm_exec.js');

(async () => {
  const go = new Go();
  const bytes = fs.readFileSync(path.join(__dirname, 'isvalid.wasm'));
  const { instance } = await WebAssembly.instantiate(bytes, go.importObject);
  go.run(instance);

  const validExpr = 'products.botd.data.bot.result == "notDetected"';
  const invalidExpr = 'products.identification.data.nonExistentField == true';

  console.log('valid:', global.isExpressionValid(validExpr));
  console.log('invalid:', global.isExpressionValid(invalidExpr));
})();
