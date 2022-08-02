package main

import (
	"os"
	"fmt"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func getPassFile(passName string) string {
	bpHome := os.Getenv("BP_HOME")

	if bpHome == "" {
		bpHome = fmt.Sprintf("%v/.local/share/bpass", os.Getenv("HOME"))
	}

	return fmt.Sprintf("%v/%v", bpHome, passName)
}

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func userError() {
	printError("Usage: bpass <new|get> <password name>\nSee the README for more information.")
}
