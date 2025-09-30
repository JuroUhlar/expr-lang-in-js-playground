#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Build using syscall/js (simpler approach)
tinygo build -o "$SCRIPT_DIR/main.wasm" -target wasm "$SCRIPT_DIR/main_syscalljs.go"
