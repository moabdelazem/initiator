package projects

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Project represents a project.
type Project interface {
	Create() error
}

// ProjectSteps represents the steps required to create a project.
type ProjectSteps struct {
	Name    string
	Action  func() error
	Message string
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
		return &NodeProject{
			Name:        name,
			Dir:         dir,
			ProjectType: "", // Will prompt user during creation
		}
	case GoLang:
		return &GoProject{
			Name:        name,
			Dir:         dir,
			ProjectType: "", // Will prompt user during creation
		}
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

type projectOption struct {
	Type        ProjectType
	Name        string
	Description string
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
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	options := []projectOption{
		{
			Type:        NodeJS,
			Name:        "Node.js (TypeScript)",
			Description: "Create a Node.js project with TypeScript setup",
		},
		{
			Type:        GoLang,
			Name:        "Go",
			Description: "Create a Go project with modern project structure",
		},
	}

	fmt.Printf("\n%s Select a project type:\n\n", white("📋"))

	// Print options with descriptions
	for i, opt := range options {
		fmt.Printf("%s %s\n", cyan(fmt.Sprintf("%d.", i+1)), opt.Name)
		fmt.Printf("   %s\n", yellow(opt.Description))
	}

	fmt.Printf("\n%s Enter your choice (1-%d): ", white("→"), len(options))

	var choice int
	for {
		fmt.Scanln(&choice)
		if choice >= 1 && choice <= len(options) {
			fmt.Printf("%s Selected: %s\n\n", white("✓"), cyan(options[choice-1].Name))
			return options[choice-1].Type
		}
		fmt.Printf("%s Please enter a number between 1 and %d: ", yellow("!"), len(options))
	}
}
