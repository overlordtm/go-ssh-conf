//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Bootstrap() error {
	return sh.RunV("go", "install", "github.com/spf13/cobra-cli@latest")
}

func Clean() error {
	fmt.Println("Cleaning the project...")
	return sh.RunV("rm", "ssh-conf")
}

func Build() error {
	fmt.Println("Building the project...")
	return sh.RunV("go", "build", "-buildvcs=true", "-o", "ssh-conf", "main.go")
}
