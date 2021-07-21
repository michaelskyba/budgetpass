package main

import (
	"fmt"
	"os"
	// "crypto/aes"
	// "crypto/cipher"
)

func main() {
	// Exit if no command
	// os.Args includes "bpass", so we check if 1, not 0
	if len(os.Args) == 1 {
		os.Exit(1)
	}

	args := os.Args[1]
	switch args {

	// New password
	case "new":
		password_name := ""
		master_password := ""
		local_password := ""

		if len(os.Args) >= 3 {
			// Password name included
			password_name = os.Args[2]
		} else {
			// Password name excluded
			fmt.Printf("Enter the name of this password: ")
			fmt.Scanln(&password_name)
		}

		// Get main input
		fmt.Printf("Enter your master password: ")
		fmt.Scanln(&master_password)

		fmt.Printf("Enter the password for '%v': ", password_name)
		fmt.Scanln(&local_password)

	// Command not found
	default:
		os.Exit(1)

	}
}
