//go:build tinygo

package main

import (
	"unsafe"

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

// This function is exported to JavaScript, so can be called using
// exports.isValid() in JavaScript.
// Takes a pointer to a string and its length
//
//export isValid
func isValid(ptr *byte, length int) int {
	// Convert pointer + length to string
	expression := ptrToString(ptr, length)
	
	// Note: This doesn't work with TinyGo due to reflection limitations
	// The expr library's parser fails to recognize operators
	_, err := expr.Compile(expression, expr.Env(Env{}))
	if err == nil {
		return 1 // true
	}
	return 0 // false
}

func ptrToString(ptr *byte, length int) string {
	if length == 0 {
		return ""
	}
	// Create a byte slice from the pointer
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = *(*byte)(unsafe.Add(unsafe.Pointer(ptr), i))
	}
	return string(bytes)
}

func main() {
	// Empty main - we're using exports only
}