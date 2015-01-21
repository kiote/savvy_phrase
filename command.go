package main

import (
	"fmt"

	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s \"first phrase\" \"second phrase\"\n", os.Args[0])
		os.Exit(1)
	}

	first := os.Args[1]
	second := os.Args[2]

	fmt.Printf("\"%s\" - %d results\n", first, 10000)
	fmt.Printf("\"%s\" - %d results\n", second, 100000)
	fmt.Println("Nothing to say")
}
