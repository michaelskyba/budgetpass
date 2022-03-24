package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"bufio"
)

func createPassword(passFile string, scanner *bufio.Scanner) {
	// Master password
	scanner.Scan()
	input := scanner.Text()
	
	if len(input) > 32 {
		fmt.Println("Master password must be <= 32 characters.")
		os.Exit(1)
	}

	// Add trail so that masterPassword == 32 characters
	trail := strings.Repeat("0", 32 - len(input))
	masterPassword := []byte(fmt.Sprintf("%v%v", input, trail))

	// Local password
	scanner.Scan()
	input = scanner.Text()
	localPassword := []byte(input)

	encrypted := encrypt(masterPassword, localPassword)

	err := os.WriteFile(passFile, encrypted, 0666)
	handle(err)
}

func getPassword(passFile string, scanner *bufio.Scanner) {
	// Master password
	scanner.Scan()
	input := scanner.Text()

	if len(input) > 32 {
		input = input[:32]
	}

	// Add trail so that masterPassword == 32 characters
	trail := strings.Repeat("0", 32 - len(input))
	masterPassword := []byte(fmt.Sprintf("%v%v", input, trail))

	encrypted, err := ioutil.ReadFile(passFile)
	handle(err)

	fmt.Println(string(decrypt(masterPassword, encrypted)))
}
