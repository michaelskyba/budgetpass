package main

import (
	"os"
	"fmt"
)

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
