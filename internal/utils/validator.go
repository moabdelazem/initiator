package utils

import (
	"fmt"
	"regexp"
)

// ValidateProjectName checks if the project name contains any invalid characters
func ValidateProjectName(name string) error {
	// Define regex for valid project name (alphanumeric, hyphens, and underscores)
	validNamePattern := regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)

	if !validNamePattern.MatchString(name) {
		return fmt.Errorf("\n❌ Invalid project name: '%s'\n"+
			"   Project names can only contain:\n"+
			"   - Letters (a-z, A-Z)\n"+
			"   - Numbers (0-9)\n"+
			"   - Hyphens (-)\n"+
			"   - Underscores (_)\n", name)
	}

	return nil
}
