package main

import (
	"log"
	"os/exec"
)

func main() {
	err := createFile("sample.txt")
	if err != nil {
		log.Printf("Something went wrong: %v", err)
	}
}

func createFile(name string) error {
	return exec.Command("touch", name).Run()
}
