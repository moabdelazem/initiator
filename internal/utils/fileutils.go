package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// getAbsPath returns the absolute path of the given target directory.
// It uses the `filepath.Abs` function to resolve the absolute path.
//
// Parameters:
//   - targetDir: The directory path to convert to an absolute path.
//   - projectName: The name of the project to append to the path
//
// Returns:
//   - string: The absolute path of the target directory.
//   - error: An error if the absolute path cannot be resolved.
//
// Example:
//
//	absPath, err := getAbsPath("relative/path", "project_name")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Absolute Path:", absPath)
func GetAbsPath(targetDir string, projectName string) (string, error) {
	// Get The Absolute Path of target directory
	absPath, err := filepath.Abs(targetDir)
	if err != nil {
		return "", err
	}

	// Append project name to the path
	fullPath := filepath.Join(absPath, projectName)

	return fullPath, nil
}

// CreateTheProjectDir creates the project directory at the given path.
// It uses the `os.MkdirAll` function to create the directory with the specified permissions.
// If the directory already exists, it will not return an error unless the path is not a directory.
//
// Parameters:
//   - path: The path of the directory to create.
//   - perm: The permission bits for the directory (e.g., 0755).
//   - initGit: A boolean flag to indicate whether to initialize a git repository.
//
// Returns:
//   - error: An error if the directory cannot be created or is not writable.
func CreateProjectDir(path string, perm os.FileMode, initGit bool) error {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("\n🚀 Setting up project in: %s\n\n", cyan(path))

	if path == "" {
		return fmt.Errorf("%s Path cannot be empty", red("✘"))
	}

	s := CreateSpinner("Validating project directory...")
	s.Start()
	if err := CheckIfDirExists(path); err != nil {
		s.Stop()
		return fmt.Errorf("%s Directory validation failed: %v", red("✘"), err)
	}
	s.Stop()
	fmt.Printf("%s Directory validation complete\n", green("✓"))

	s = CreateSpinner("Creating project directory...")
	s.Start()
	if err := os.MkdirAll(path, perm); err != nil {
		s.Stop()
		return fmt.Errorf("%s Failed to create directory: %v", red("✘"), err)
	}
	s.Stop()
	fmt.Printf("%s Project directory created\n", green("✓"))

	if initGit {
		if shouldInit := promptUserForGit(); shouldInit {
			if err := initializeGitRepository(path); err != nil {
				return err
			}
		}
	}

	fmt.Printf("\n%s Project setup completed successfully!\n", green("✨"))
	return nil
}

// checkIfDirExists checks if the directory exists at the given path.
// It uses the `os.Stat` function to check if the directory exists.
// Parameters:
//   - path: The path of the directory to check.
//
// Returns:
//   - error: An error if the directory exists and is not a directory.
func CheckIfDirExists(path string) error {
	if info, err := os.Stat(path); err == nil {
		if !info.IsDir() {
			return fmt.Errorf("path %s already exists and is not a directory", path)
		}

		// Directory exists, ask user what to do
		if shouldOverwrite := promptUserForOverwrite(path); !shouldOverwrite {
			return fmt.Errorf("operation cancelled by user")
		}

		// If user wants to overwrite, remove the existing directory
		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to remove existing directory: %v", err)
		}
	}
	return nil
}

// promptUserForOverwrite asks the user whether they want to overwrite an existing directory.
// It prompts the user for a yes/no response and handles the input validation.
//
// Parameters:
//   - path: The path to the directory that already exists
//
// Returns:
//   - bool: true if user confirms overwrite, false otherwise
//
// The function will continuously prompt until a valid response is received.
// Valid responses are:
//   - "y" or "yes" (case insensitive) for confirmation
//   - "n", "no", or empty string (case insensitive) for rejection
func promptUserForOverwrite(path string) bool {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	fmt.Printf("\n%s Directory Already Exists\n", yellow("⚠"))
	fmt.Printf("%s Location: %s\n", white("→"), cyan(path))
	fmt.Printf("\nOptions:\n")
	fmt.Printf("  %s Remove existing and create new\n", cyan("y"))
	fmt.Printf("  %s Cancel operation\n", cyan("n"))

	fmt.Printf("\n%s Your choice [y/N]: ", white("→"))

	reader := bufio.NewReader(os.Stdin)
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s Error reading input\n", yellow("!"))
			return false
		}

		response = strings.ToLower(strings.TrimSpace(response))
		switch response {
		case "y", "yes":
			fmt.Printf("%s Proceeding with overwrite\n", white("✓"))
			return true
		case "", "n", "no":
			fmt.Printf("%s Operation cancelled\n", yellow("✗"))
			return false
		default:
			fmt.Printf("%s Please answer with 'y' or 'n': ", yellow("!"))
		}
	}
}

// InitGitRepo initializes a new Git repository in the specified directory.
// It runs the 'git init' command in the given directory path.
//
// Parameters:
//   - dir: The directory path where the Git repository should be initialized
//
// Returns:
//   - error: Returns nil on success, or an error if the git init command fails
//     with the command output appended to the error message
func InitGitRepo(dir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git init failed: %v - %s", err, string(output))
	}

	log.Println("Initialized new git repository")
	return nil
}

// promptUserForGit prompts the user to decide whether to initialize a git repository.
// It reads user input from stdin and expects a yes/no answer.
// The function keeps prompting until a valid response is received.
//
// Valid responses (case-insensitive):
// - Yes: "y", "yes", or empty (press enter)
// - No: "n", "no"
//
// Returns:
//   - true if user wants to initialize git repository
//   - false if user declines or if there's an error reading input
func promptUserForGit() bool {
	cyan := color.New(color.FgCyan).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("\n%s Git Repository Setup\n", white("📦"))
	fmt.Printf("\nWould you like to:\n")
	fmt.Printf("  %s Initialize a new Git repository\n", cyan("y"))
	fmt.Printf("  %s Skip Git initialization\n", cyan("n"))

	fmt.Printf("\n%s Your choice [Y/n]: ", white("→"))

	reader := bufio.NewReader(os.Stdin)
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s Error reading input\n", yellow("!"))
			return false
		}

		response = strings.ToLower(strings.TrimSpace(response))
		switch response {
		case "", "y", "yes":
			fmt.Printf("%s Initializing Git repository\n", white("✓"))
			return true
		case "n", "no":
			fmt.Printf("%s Skipping Git initialization\n", white("→"))
			return false
		default:
			fmt.Printf("%s Please answer with 'y' or 'n': ", yellow("!"))
		}
	}
}

func initializeGitRepository(path string) error {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// Initialize git repository
	s := CreateSpinner("Initializing Git repository...")
	s.Start()
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		s.Stop()
		return fmt.Errorf("%s Git initialization failed: %v", red("✘"), err)
	}
	s.Stop()
	fmt.Printf("%s Git repository initialized\n", green("✓"))

	// Create .gitignore file
	s = CreateSpinner("Creating .gitignore file...")
	s.Start()
	gitignore := `# Dependencies
node_modules/
npm-debug.log*
yarn-debug.log*
yarn-error.log*

# Production build
dist/
build/

# Environment variables
.env
.env.local
.env.*.local

# IDE/Editor specific
.idea/
.vscode/
*.swp
*.swo

# OS specific
.DS_Store
Thumbs.db`

	if err := os.WriteFile(filepath.Join(path, ".gitignore"), []byte(gitignore), 0644); err != nil {
		s.Stop()
		return fmt.Errorf("%s Failed to create .gitignore: %v", red("✘"), err)
	}
	s.Stop()
	fmt.Printf("%s Created .gitignore file\n", green("✓"))

	return nil
}
