package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	version := "0.1.1"
	stable := false

	fmt.Printf("Welcome to BudgetPass version %v. To say that it's stable would be %v.\n", version, stable)
	fmt.Printf("Enter a command (help): ")

	// Get input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	switch input {
	case "help":
		fmt.Println("currently implemented commands:")
		fmt.Println("- help")
	default:
		fmt.Printf("command not found: %v\n", input)
	}
}
