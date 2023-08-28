//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	moduleName string = "main"
	binaryName string = "mahgraphql"
)

var (
	binaryOutputDir string = "." + string(os.PathSeparator)
	binaryOutput           = binaryOutputDir + binaryName
	entryGoFile            = path.Join("main.go")
)

// Run installs 3rd party libraries, run unit tests, run swagger, compile binary in one go.
// Usually recommended for first setup or prepping for PRs.
func Setup() {
	mg.Deps(Test, Build)
}

// Compiles developer executable on their host machine.
func Build() error {
	fmt.Println("Go generate...")
	err := sh.RunV("go", "generate", "./...")
	if err != nil {
		return err
	}

	fmt.Println("Building...")
	return sh.RunV("go", "build", "-o", binaryName, binaryOutputDir)
}

// Run main
func Run() error {
	fmt.Printf("Running %s...\n", binaryOutput)
	return sh.RunV(binaryOutput)
}

// Run Go and Ginkgo tests.
func Test() error {
	return sh.RunV("go", "test", "-v", "./...")
}

// Cleanup the code and display any typos and or linting errors.
func Lint() error {
	fmt.Println("Running go fmt")
	err := sh.Run("go", "fmt")

	if err != nil {
		return err
	}

	fmt.Println("Checking for typos")
	err = sh.RunV("misspell", ".")

	if err != nil {
		return err
	}

	fmt.Println("Checking for linting errors")
	return sh.RunV("golangci-lint", "run")
}

// Removes old binaries.
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll(binaryOutput)
}
