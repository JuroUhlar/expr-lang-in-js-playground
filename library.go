package main

import (
	"github.com/expr-lang/expr"
)

// Environment structs --------------------------------

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

// ----------------------------------------------------------------------------

func isExpressionValid(expression string) bool {
	_, err := expr.Compile(expression, expr.Env(Env{}))
	return err == nil
}

