package main

import (
	"os"
	"bufio"
)

func main() {
	// os.Args includes "bpass" (bpass command password), so we check if < 3, not 2
	if len(os.Args) < 3 {
		userError()
	}

	command := os.Args[1]
	passFile := getPassFile(os.Args[2])
	scanner := bufio.NewScanner(os.Stdin)

	switch command {

	case "new":
		createPassword(passFile, scanner)

	case "get":
		getPassword(passFile, scanner)

	default:
		userError()
	}
}
