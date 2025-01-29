package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
//
// Returns:
//   - error: An error if the directory cannot be created or is not writable.
func CreateProjectDir(path string, perm os.FileMode) error {
	// Validate input
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	// Check if the directory already exists
	if err := CheckIfDirExists(path); err != nil {
		return err
	}

	// Create the directory
	if err := os.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("failed to create directory at %s: %v", path, err)
	}

	// Verfiy the directory is writeable
	testFile := filepath.Join(path, ".test")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return fmt.Errorf("failed to write to directory %s: %v", path, err)
	}
	// Remove the test file
	os.Remove(testFile)

	// Log the success
	log.Printf("Created project directory at %s", path)

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
	// Prompt the user for confirmation
	reader := bufio.NewReader(os.Stdin)

	// Loop until a valid response is received
	for {
		fmt.Printf("Directory '%s' already exists. Do you want to overwrite it? [y/N]: ", path)
		// Read the user's response
		response, err := reader.ReadString('\n')
		// Handle any errors reading the input
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return false
		}

		// Normalize the response and check for valid values
		response = strings.ToLower(strings.TrimSpace(response))
		switch response {
		case "y", "yes":
			return true
		case "", "n", "no":
			return false
		default:
			fmt.Println("Please answer with 'y' or 'n'")
		}
	}
}
