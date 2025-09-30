//go:build js && wasm

package main

import "syscall/js"

// isExpressionsValidJS is a wrapper for the isExpressionValid function that can be called from JavaScript
func isExpressionsValidJS(_ js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return false
	}
	return isExpressionValid(args[0].String())
}

func main() {
	js.Global().Set("isExpressionValid", js.FuncOf(isExpressionsValidJS))

	// Blocks forever, keeping the WASM module alive so JavaScript can call the functions
	<-make(chan struct{})
}
