package main

import (
	"fmt"

	"os"
)

func main() {
	args := os.Args[1:]

	first := args[0]
	second := args[1]

	fmt.Printf("\"%s\" - %d results\n", first, 10000)
	fmt.Printf("\"%s\" - %d results\n", second, 100000)
	fmt.Println("Nothing to say")
}
