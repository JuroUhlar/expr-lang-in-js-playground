//go:build !js || !wasm

package main

func main() {
	// no-op for native build; tests use exported functions directly
}
