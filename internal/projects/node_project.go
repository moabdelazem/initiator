package projects

import (
	"fmt"
	"os/exec"
)

// NodeProject represents a Node.js project.
type NodeProject struct {
	Name string
	Dir  string
}

// Create initializes a new Node.js project in the specified directory.
// It first changes to the project directory, then runs 'npm init -y' to create
// a new Node.js project with default settings.
// Returns an error if directory change fails or if npm initialization fails.
func (p *NodeProject) Create() error {
	// Navigate to the project directory
	if err := ChangeDirectory(p.Dir); err != nil {
		return err
	}

	// Initialize Node.js project
	cmd := exec.Command("npm", "init", "-y")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Node.js project: %v", err)
	}

	fmt.Printf("Node.js project '%s' created successfully\n", p.Name)
	return nil
}
