package main

import (
	"fmt"
	"strings"
	"os"
	"crypto/aes"
	"crypto/cipher"
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
		input := ""

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
		fmt.Scanln(&input)
		
		// Make sure master_password <= 32 characters
		if len(input) > 32 {
			fmt.Println("Master password must be <= 32 characters.")
			os.Exit(1)
		}

		// Add trail so that master_password == 32 characters
		trail := strings.Repeat("0", 32 - len(input))
		master_password := []byte(fmt.Sprintf("%v%v", input, trail))

		fmt.Printf("Enter the password for '%v': ", password_name)
		fmt.Scanln(&input)
		local_password := []byte(input)

		// Encryption
		my_cipher, _ := aes.NewCipher(master_password)
		gcm, _ := cipher.NewGCM(my_cipher)
		nonce := make([]byte, gcm.NonceSize())
		encrypted := gcm.Seal(nonce, nonce, local_password, nil)

		// Print to file
		// message := []byte(fmt.Sprintf("%v: encrypt %v using %v\n", password_name, local_password, master_password))
		err := os.WriteFile(password_name, encrypted, 0666)
		if err != nil {
			fmt.Println(err)
		}

	// Command not found
	default:
		os.Exit(1)

	}
}
