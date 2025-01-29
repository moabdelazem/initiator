package utils

import "fmt"

// enum for the supported languages
const (
	_ = iota + 1
	Node
	GoLang
)

// Project represents the project structure
type Project struct {
	Name     string
	Language int
}

// NewProject creates a new Project instance.
// It validates the provided language and returns an error if the language is unsupported.
//
// Parameters:
//   - name: The name of the project.
//   - lang: The programming language for the project.
//
// Returns:
//   - *Project: A pointer to the created Project instance.
//   - error: An error if the language is unsupported.
func NewProject(name string, lang int) *Project {
	if lang != Node && lang != GoLang {
		fmt.Println("Invalid Language")
		return nil
	}
	return &Project{
		Name:     name,
		Language: lang,
	}
}

// Create initializes the project based on its language.
//
// Returns:
//   - error: An error if the project creation fails.
func (p *Project) Create() error {
	switch p.Language {
	case Node:
		return createNodeProject(p)
	case GoLang:
		return createGoProject(p)
	default:
		return fmt.Errorf("unsupported language")
	}
}
