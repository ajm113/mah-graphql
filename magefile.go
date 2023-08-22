//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

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
	err := sh.RunV("go", "generate")
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
	fmt.Println("Running go tests")

	// Run stripe-mock
	stripeMock := exec.Command("stripe-mock", "-http-port", "12111", "-https-port", "12112")
	stripeMock.Stderr = os.Stderr
	stripeMock.Stdout = os.Stdout
	stripeMock.Start()

	// Give stripe-mock a tiny bit of time to start.
	time.Sleep(time.Millisecond * 200)

	err := sh.RunV("go", "test", "-v", "./...")

	if err := stripeMock.Process.Kill(); err != nil {
		log.Fatal("failed to kill process: ", err)
	}

	return err
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
