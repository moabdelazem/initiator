package projects

import (
	"fmt"
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
	return nil
}
