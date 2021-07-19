package main

import (
	"fmt"
	"os"
)

func main() {
	// Exit if no command
	// os.Args includes "bpass", so we check if 1, not 0
	if len(os.Args) == 1 {
		fmt.Println("see 'bpass help' for usage instructions")
		os.Exit(1)
	}

	args := os.Args[1]
	switch args {
	case "help":
		fmt.Println("currently implemented commands:")
		fmt.Println("- help")
	default:
		fmt.Printf("command not found: %v\n", args)
	}
}
