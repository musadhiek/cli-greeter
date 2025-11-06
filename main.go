package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	// ExitSuccess indicates successful execution
	ExitSuccess = 0
	// ExitError indicates an error occurred
	ExitError = 1
)

func main() {
	if err := run(flag.CommandLine, os.Args[1:]); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(ExitError)
	}
	os.Exit(ExitSuccess)
}

// run executes the main application logic
func run(fs *flag.FlagSet, args []string) error {
	name, err := greet(fs, args)
	if err != nil {
		return fmt.Errorf("failed to greet: %w", err)
	}
	fmt.Println(name)
	return nil
}

// greet parses command-line flags and returns the name to greet
func greet(fs *flag.FlagSet, args []string) (string, error) {
	name := fs.String("name", "World", "Name to greet")

	if err := fs.Parse(args); err != nil {
		return "", fmt.Errorf("failed to parse flags: %w", err)
	}

	// Environment variable takes precedence over default, but flag takes precedence over env
	if *name == "World" {
		if envName := os.Getenv("GREET_NAME"); envName != "" {
			*name = envName
		}
	}

	return *name, nil
}
