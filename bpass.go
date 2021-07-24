package main

import (
	"fmt"
	"strings"
	"os"
	"crypto/aes"
	"crypto/cipher"
)

func main() {
	usage := "See the README for usage."

	// os.Args includes "bpass" (bpass command password), so we check if < 3, not 2
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	command := os.Args[1]
	password_name := os.Args[2]

	switch command {

	// New password
	case "new":
		input := ""

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
		err := os.WriteFile(password_name, encrypted, 0666)
		if err != nil {
			fmt.Println(err)
		}

	// Command not found
	default:
		fmt.Println(usage)
		os.Exit(1)

	}
}
