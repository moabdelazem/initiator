package projects

import (
	"os"
	"testing"
)

// TestHelperProcess isn't a real test. It's used as a helper process for TestGoProject_Create.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	// Get the command arguments
	args := os.Args
	for i, arg := range args {
		if arg == "--" {
			args = args[i+1:]
			break
		}
	}

	// Mock different commands based on the command name and arguments
	switch args[0] {
	case "go":
		if len(args) > 1 && args[1] == "mod" && args[2] == "init" {
			// Simulate successful go mod init
			os.Exit(0)
		}
	}

	// Default: command not handled
	os.Exit(1)
}
