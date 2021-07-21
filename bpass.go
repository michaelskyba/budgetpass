package main

import (
	"fmt"
	"os"
	// "crypto/aes"
	// "crypto/cipher"
	// "io/ioutil"
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
		fmt.Printf("Enter your master password: ")
		fmt.Scanln(&master_password)

		fmt.Printf("Enter the password for '%v': ", password_name)
		fmt.Scanln(&local_password)

	// Help
	case "help":
		fmt.Println("currently implemented commands:")
		fmt.Println("- help")
		fmt.Println("- new \"[password name]\"")

	// Command not found
	default:
		fmt.Printf("command not found: %v\n", args)

	}
}
