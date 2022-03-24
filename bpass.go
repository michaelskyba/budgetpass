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

func getPassFile(passName string) string {
	bpHome := os.Getenv("BP_HOME")

	if bpHome == "" {
		bpHome = fmt.Sprintf("%v/.local/share/bpass", os.Getenv("HOME"))
	}

	return fmt.Sprintf("%v/%v", bpHome, passName)
}

func userError() {
	fmt.Println("Usage: bpass <new|get> <password name>\nSee the README for more information.")
	os.Exit(1)
}

func main() {
	// os.Args includes "bpass" (bpass command password), so we check if < 3, not 2
	if len(os.Args) < 3 {
		userError()
	}

	command := os.Args[1]
	passFile := getPassFile(os.Args[2])

	var input string

	switch command {

	case "new":
		// Get main input
		fmt.Printf("Enter your master password: ")
		fmt.Scanln(&input)
		
		// Make sure masterPassword <= 32 characters
		if len(input) > 32 {
			fmt.Println("Master password must be <= 32 characters.")
			os.Exit(1)
		}

		// Add trail so that masterPassword == 32 characters
		trail := strings.Repeat("0", 32 - len(input))
		masterPassword := []byte(fmt.Sprintf("%v%v", input, trail))

		fmt.Printf("Enter the password for '%v': ", passFile)
		fmt.Scanln(&input)
		localPassword := []byte(input)

		// Encryption
		myCipher, err := aes.NewCipher(masterPassword)
		handle(err)

		gcm, err := cipher.NewGCM(myCipher)
		handle(err)

		nonce := make([]byte, gcm.NonceSize())
		encrypted := gcm.Seal(nonce, nonce, localPassword, nil)

		// Print to file
		err = os.WriteFile(passFile, encrypted, 0666)
		handle(err)

	case "get":
		// Get main input
		fmt.Printf("Enter your master password: ")
		fmt.Scanln(&input)

		// Make sure masterPassword <= 32 characters
		if len(input) > 32 {
			input = input[:32]
		}

		// Add trail so that masterPassword == 32 characters
		trail := strings.Repeat("0", 32 - len(input))
		masterPassword := []byte(fmt.Sprintf("%v%v", input, trail))

		// Get contents of password file
		encrypted, err := ioutil.ReadFile(passFile)
		handle(err)

		// Decrypt
		myCipher, err := aes.NewCipher(masterPassword)
		handle(err)

		gcm, err := cipher.NewGCM(myCipher)
		handle(err)

		nonceSize := gcm.NonceSize()
		nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]

		localPassword, err := gcm.Open(nil, nonce, encrypted, nil)
		handle(err)

		// Done - output
		fmt.Println(string(localPassword))

	// Command not found
	default:
		userError()
	}
}
