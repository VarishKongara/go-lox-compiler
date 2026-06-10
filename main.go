package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println("Usage: [executable] [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		// runFile(args[0])
	} else {
		// runPrompt()
	}
}

func runFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	run(string(data))
	// Indicates an error using the exit code
	if hadError {
		os.Exit(65)
	}
	return nil
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = strings.TrimRight(line, "\r\n")
		run(line)
		hadError = false // reset without killing session
	}
	return nil
}

func run(source string) error {
	//Scanner has not been created yet
	scanner := Scanner{Source: source}
	tokens, err := scanner.ScanTokens()
	if err != nil {
		return err
	}
	for _, token := range tokens {
		fmt.Println(token)
	}

	return nil
}

var hadError bool = false

func createError(line int, message string) {
	reportError(line, "", message)
}

func reportError(line int, where string, message string) {
	fmt.Fprintln(os.Stdout, "[line %v] Error %v: %v", line, where, message)
	hadError = true
}
