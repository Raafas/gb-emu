package main

import (
	"fmt"
	"gb-emu/cart"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_cart_file>")
		return
	}

	cartFile := os.Args[1]
	ctx := cart.CartContext{}

	if ctx.CartLoad(cartFile) {
		fmt.Println("Cartridge loaded successfully.")
	} else {
		fmt.Println("Failed to load cartridge.")
	}
}
