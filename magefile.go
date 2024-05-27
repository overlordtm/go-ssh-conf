//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Clean() error {
	fmt.Println("Cleaning the project...")
	return sh.RunV("rm", "ssh-conf")
}

func Build() error {
	fmt.Println("Building the project...")
	return sh.RunV("go", "build", "-o", "ssh-conf", "main.go")
}
