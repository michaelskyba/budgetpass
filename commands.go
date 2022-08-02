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
		printError("Master password must be <= 32 characters.")
	}

	// Add trail so that masterPassword == 32 characters
	trail := strings.Repeat("0", 32 - len(input))
	masterPassword := []byte(fmt.Sprintf("%v%v", input, trail))

	scanner.Scan()
	localPassword := []byte(scanner.Text())

	encrypted := encrypt(masterPassword, localPassword)

	err := os.WriteFile(passFile, encrypted, 0666)
	hdl(err)
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
	hdl(err)

	fmt.Println(string(decrypt(masterPassword, encrypted)))
}
