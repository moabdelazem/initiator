package projects

import (
	"fmt"
	"os"
)

// Project represents a project.
type Project interface {
	Create() error
}

// ProjectType represents the type of project.
type ProjectType string

const (
	NodeJS ProjectType = "nodejs"
	GoLang ProjectType = "golang"
)

// NewProject creates and returns a new Project instance based on the specified project type.
// It takes the project name, directory path, and project type as parameters.
// For NodeJS projects, it returns a NodeProject instance.
// For GoLang projects, it returns a GoProject instance.
// Returns nil for unsupported project types.
func NewProject(name string, dir string, projectType ProjectType) Project {
	switch projectType {
	case NodeJS:
		return &NodeProject{Name: name, Dir: dir}
	case GoLang:
		return &GoProject{Name: name, Dir: dir}
	default:
		return nil
	}
}

// ChangeDirectory changes the current working directory to the specified directory path.
// It returns an error if the directory change operation fails.
//
// Parameters:
//   - dir: The target directory path to change to
//
// Returns:
//   - error: nil if successful, otherwise returns an error with details about the failure
func ChangeDirectory(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", dir, err)
	}
	return nil
}
