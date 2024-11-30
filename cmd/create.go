package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create [mod_name]",
	Short: "Create a new project",
	Long:  `Create a new Go project with the specified module name and target path.`,
	Args:  cobra.ExactArgs(1), // Require exactly one argument: mod_name
	Run: func(cmd *cobra.Command, args []string) {
		modName := args[0]

		// Validate module name
		if err := validateModuleName(modName); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Get the target path from the flag
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Printf("Error retrieving path: %v\n", err)
			return
		}

		// Convert the path to an absolute path
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Create the project directory
		if err := initProject(absPath); err != nil {
			fmt.Printf("Error initializing project: %v\n", err)
			return
		}

		// Initialize the Go module
		if err := initGoMod(modName, absPath); err != nil {
			fmt.Printf("Error initializing Go module: %v\n", err)
			return
		}

		// Create base files
		if err := createBaseFiles(absPath); err != nil {
			fmt.Printf("Error creating base files: %v\n", err)
			return
		}

		fmt.Printf("Project created successfully at %s\n", absPath)
	},
}

// Validate the module name
func validateModuleName(modName string) error {
	regex := `^[a-zA-Z0-9\.\-/]+$`
	matched, err := regexp.MatchString(regex, modName)
	if err != nil {
		return fmt.Errorf("failed to validate module name: %v", err)
	}
	if !matched {
		return fmt.Errorf("invalid module name: must contain only alphanumeric characters, periods, hyphens, or slashes")
	}
	return nil
}

// Create the project directory
func initProject(projectPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists. Overwrite? [y/N]: ", projectPath)
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
			return fmt.Errorf("project creation canceled")
		}

		// Delete existing directory
		if err := os.RemoveAll(projectPath); err != nil {
			return fmt.Errorf("failed to remove existing directory: %v", err)
		}
	}

	// Create the directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}

// Initialize a Go module
func initGoMod(modName, projectPath string) error {
	// Navigate to the project directory
	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %v", err)
	}

	// Execute `go mod init`
	cmd := exec.Command("go", "mod", "init", modName)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Go module: %v", err)
	}

	return nil
}

func createBaseFiles(absPath string) error {
	mainGoPath := filepath.Join(absPath, "cmd", "main.go")
	if err := os.MkdirAll(filepath.Dir(mainGoPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, From Initiator!")
}
	`

	if err := os.WriteFile(mainGoPath, []byte(mainGoContent), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	makefilePath := filepath.Join(absPath, "Makefile")
	makefileContent := `build:
	go build -o bin/main ./cmd

run:
	go run ./cmd/main.go

clean:
	rm -rf bin
`

	if err := os.WriteFile(makefilePath, []byte(makefileContent), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	// Creatre .gitignore file
	gitignorePath := filepath.Join(absPath, ".gitignore")
	gitignoreContent := `bin/
*.log
*.out
*.exe
`

	if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	if err := initGitRepo(absPath); err != nil {
		return fmt.Errorf("failed to initialize Git repository: %v", err)
	}

	return nil
}

func initGitRepo(absPath string) error {
	cmd := exec.Command("git", "init")
	// Change the branch name to master by default
	cmd.Env = append(os.Environ(), "GIT_INIT_DEFAULT_BRANCH=master")
	cmd.Dir = absPath
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Git repository: %v", err)
	}
	fmt.Println("Git repository initialized")
	return nil
}
