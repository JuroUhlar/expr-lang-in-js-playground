## Using expr-lang in JavaScript environments (POC/Playground)

Expr-lang is a Go-centric expression language. You might have chosen for your Go server and then discover you need to use it in the browser or another JavaScript environment.

This repo is to find the best way.



## Run

```bash
# Navigate to go directory
cd go

# Run directly
go run helloworld.go

# Or build and run
go build -o helloworld helloworld.go
./helloworld
```

## WASM build and Node usage

Build the validator to WebAssembly and run it from Node.

### Build

From the `go` directory:

```bash
cd go
mkdir -p wasm
GOOS=js GOARCH=wasm go build -o ./wasm/isvalid.wasm .
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./wasm/
```

### Run in Node

Create `go/wasm/test.js`:

```javascript
const fs = require('fs');
const path = require('path');
require('./wasm_exec.js');

(async () => {
  const go = new Go();
  const ready = new Promise((resolve) => {
    global.__registerExprExports = (exportsObj) => {
      global.__exprExports = exportsObj;
      resolve();
    };
  });
  const bytes = fs.readFileSync(path.join(__dirname, 'isvalid.wasm'));
  const { instance } = await WebAssembly.instantiate(bytes, go.importObject);
  go.run(instance);
  await ready;

  const validExpr = 'products.botd.data.bot.result == "notDetected"';
  const invalidExpr = 'products.identification.data.nonExistentField == true';

  console.log('valid:', global.__exprExports.isExpressionValid(validExpr));
  console.log('invalid:', global.__exprExports.isExpressionValid(invalidExpr));
})();
```

Run it:

```bash
node go/wasm/test.js
```

### Notes

- Export API: `isExpressionValid(string): boolean` provided via a registration callback; no globals required.
- Unknown nested fields are rejected at compile time by expr’s environment typing.
- Docs: Environment — https://expr-lang.org/docs/environment


