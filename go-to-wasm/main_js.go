//go:build js && wasm

package main

import (
	"syscall/js"
)

func isExpressionValidJS(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return js.ValueOf(false)
	}
	expr := args[0].String()
	return js.ValueOf(isExpressionValid(expr))
}

func main() {
	js.Global().Set("isExpressionValid", js.FuncOf(func(this js.Value, args []js.Value) any {
		return isExpressionValidJS(this, args)
	}))
	select {}
}
