package projects

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

// PromptUserForProjectType prompts the user to select a project type from predefined options.
// It continuously asks for input until a valid project type is selected.
// The function reads user input from standard input and converts it to lowercase for comparison.
// Returns:
//   - ProjectType: The selected project type (either NodeJS or GoLang)
//   - If an error occurs while reading input, returns an empty ProjectType
//
// Valid inputs are:
//   - "nodejs" for NodeJS projects
//   - "golang" for GoLang projects
func PromptUserForProjectType() ProjectType {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Select Project Type (nodejs/golang): ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return ""
		}

		response = strings.ToLower(strings.TrimSpace(response))
		switch response {
		case "nodejs":
			return NodeJS
		case "golang":
			return GoLang
		default:
			fmt.Println("Please select a valid project type (nodejs/golang)")
		}
	}
}
