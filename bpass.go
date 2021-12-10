package main

import (
	"fmt"
	"strings"
	"os"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
)

// Handle errors more elegantly
func handle(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	usage := `Usage: bpass <new|get> <password name>
See the README for more information.`
	var input string

	// os.Args includes "bpass" (bpass command password), so we check if < 3, not 2
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	command := os.Args[1]
	password_name := os.Args[2]

	// Decide where password files will be stored
	var home_dir string
	if os.Getenv("BP_HOME") == "" {
		home_dir = fmt.Sprintf("%v/.local/share/bpass", os.Getenv("HOME"))
	} else {
		home_dir = os.Getenv("BP_HOME")
	}
	password_name = fmt.Sprintf("%v/%v", home_dir, password_name)

	switch command {

	case "new":
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
		my_cipher, err := aes.NewCipher(master_password)
		handle(err)

		gcm, err := cipher.NewGCM(my_cipher)
		handle(err)

		nonce := make([]byte, gcm.NonceSize())
		encrypted := gcm.Seal(nonce, nonce, local_password, nil)

		// Print to file
		err = os.WriteFile(password_name, encrypted, 0666)
		handle(err)

	case "get":
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
		encrypted, err := ioutil.ReadFile(password_name)
		handle(err)

		// Decrypt
		my_cipher, err := aes.NewCipher(master_password)
		handle(err)

		gcm, err := cipher.NewGCM(my_cipher)
		handle(err)

		nonce_size := gcm.NonceSize()
		nonce, encrypted := encrypted[:nonce_size], encrypted[nonce_size:]

		local_password, err := gcm.Open(nil, nonce, encrypted, nil)
		handle(err)

		// Done - output
		fmt.Println(string(local_password))

	// Command not found
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
