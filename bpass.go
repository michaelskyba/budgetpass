package main

import (
	"fmt"
	"strings"
	"os"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
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

	case "get":
		input := ""

		// Get main input
		fmt.Printf("Enter your master password: ")
		fmt.Scanln(&input)

		// Make sure master_password <= 32 characters
		if len(input) > 32 {
			input = input[:32]
		}

		// Add trail so that master_password == 32 characters
		trail := strings.Repeat("0", 32 - len(input))
		master_password := []byte(fmt.Sprintf("%v%v", input, trail))

		// Get contents of password file
		encrypted, _ := ioutil.ReadFile(password_name)

		// Decrypt
		my_cipher, _ := aes.NewCipher(master_password)
		gcm, _ := cipher.NewGCM(my_cipher)
		nonce_size := gcm.NonceSize()
		nonce, encrypted := encrypted[:nonce_size], encrypted[nonce_size:]
		local_password, _ := gcm.Open(nil, nonce, encrypted, nil)
		fmt.Println(string(local_password))

	// Command not found
	default:
		fmt.Println(usage)
		os.Exit(1)

	}
}
