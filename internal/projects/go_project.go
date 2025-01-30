package projects

import (
	"fmt"
	"os"
	"os/exec"
)

// GoProject represents a Go project.
type GoProject struct {
	Name string
	Dir  string
}

// Create initializes a new Go project in the specified directory.
// It changes to the project directory and runs 'go mod init' with the project name.
// Returns an error if directory change fails or if project initialization fails.
func (p *GoProject) Create() error {
	// Navigate to the project directory
	if err := ChangeDirectory(p.Dir); err != nil {
		return err
	}

	// Initialize Go project
	cmd := exec.Command("go", "mod", "init", p.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Go project: %v", err)
	}

	fmt.Printf("Go project '%s' created successfully\n", p.Name)

	// initalize the project with a main.go file
	if err := p.initMainFile(); err != nil {
		return err
	}
	return nil
}

// initMainFile creates a main.go file in the project directory.
// It writes a simple Go program to the file.
// Returns an error if file creation or writing fails.
func (p *GoProject) initMainFile() error {
	// Create the cmd directory
	if err := os.Mkdir("cmd", 0755); err != nil {
		return fmt.Errorf("failed to create cmd directory: %v", err)
	}

	// Change to the cmd directory
	if err := ChangeDirectory("cmd"); err != nil {
		return err
	}

	// Create the main.go file
	file, err := os.Create("main.go")
	if err != nil {
		return fmt.Errorf("failed to create main.go file: %v", err)
	}
	defer file.Close()

	// Write the Go program to the file
	_, err = file.WriteString(`package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
`)
	if err != nil {
		return fmt.Errorf("failed to write to main.go file: %v", err)
	}

	fmt.Println("main.go file created successfully")
	return nil
}
