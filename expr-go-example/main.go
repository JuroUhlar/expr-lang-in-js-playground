//go:build !js || !wasm

package main

import "fmt"

var exampleExpressions = []string{
	`proxy.type == "residential" && proxy.notExistentField > 0.9`,
	`bot == true || country == "US"`,
}

func main() {
	for _, expr := range exampleExpressions {
		fmt.Println("Expression:", expr)
		fmt.Println("Is valid:", isExpressionValid(expr))
		fmt.Println()
	}
}
