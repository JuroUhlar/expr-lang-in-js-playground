#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Build the TinyGo-specific main file with explicit exports
# Use -scheduler=none to avoid runtime complexity
tinygo build -o "$SCRIPT_DIR/main.wasm" -target wasm -scheduler=none "$SCRIPT_DIR/main_tinygo.go"
