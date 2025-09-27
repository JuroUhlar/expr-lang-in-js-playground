package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func expr_demo() {
	// Basic expression evaluation
	result, err := expr.Eval("2 + 3 * 4", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("2 + 3 * 4 = %v\n", result)

	// Expression with variables
	env := map[string]interface{}{
		"name": "World",
		"age":  30,
		"pi":   3.14159,
	}

	greeting, err := expr.Eval(`"Hello, " + name + "! You are " + string(age) + " years old."`, env)
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Boolean expression
	canVote, err := expr.Eval("age >= 18", env)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Can vote: %v\n", canVote)

	// Compile and reuse expression
	program, err := expr.Compile("pi * age", expr.Env(env))
	if err != nil {
		panic(err)
	}
	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pi * age = %.2f\n", output)
}
