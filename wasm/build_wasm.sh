#!/bin/bash

# Just so you can run this from any directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$(dirname "$SCRIPT_DIR")"

# Main thing
GOOS=js GOARCH=wasm go build -o "$SCRIPT_DIR/main.wasm" .
