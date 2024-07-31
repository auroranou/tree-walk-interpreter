package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/auroranou/tree-walk-interpreter/scan"
)

var hadError bool = false

func main() {
	var filePath string

	flag.StringVar(&filePath, "file", "", "File path of program to run")
	flag.Parse()

	if filePath != "" {
		runFile(filePath)
	} else {
		runPrompt()
	}
}

func runFile(filePath string) {
	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	run(string(fileContent))

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()
		run(line)
		hadError = false
	}
}

func run(source string) {
	scanner := scan.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Printf("%v", token)
	}
}
