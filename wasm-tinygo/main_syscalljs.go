//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/expr-lang/expr"
)

type Env struct {
	Bot          bool   `expr:"bot"`
	Country      string `expr:"country"`
	Vpn          bool   `expr:"vpn"`
	SuspectScore int    `expr:"suspectScore"`
	Proxy        Proxy  `expr:"proxy"`
}

type Proxy struct {
	Type       string  `expr:"type"`
	Confidence float64 `expr:"confidence"`
}

func isExpressionValid(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return false
	}
	
	expression := args[0].String()
	_, err := expr.Compile(expression, expr.Env(Env{}))
	return err == nil
}

func main() {
	js.Global().Set("isExpressionValid", js.FuncOf(isExpressionValid))
	<-make(chan struct{})
}
