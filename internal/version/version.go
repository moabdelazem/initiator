package version

import "fmt"

var (
	// Version is the current version of the application
	Version = "0.1.0"

	// Commit is the git commit hash of the build
	Commit = "unknown"

	// BuildDate is the date when the binary was built
	BuildDate = "unknown"
)

// GetVersion returns the full version information as a string
func GetVersion() string {
	return fmt.Sprintf("Version: %s\nCommit: %s\nBuild Date: %s",
		Version, Commit, BuildDate)
}

// GetShortVersion returns just the version number
func GetShortVersion() string {
	return Version
}
