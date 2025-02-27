package version

import (
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Save original values
	origVersion := Version
	origCommit := Commit
	origBuildDate := BuildDate

	// Set test values
	Version = "1.0.0"
	Commit = "abc123"
	BuildDate = "2023-07-01"

	// Get version string
	versionStr := GetVersion()

	// Check that it contains all the expected information
	if !strings.Contains(versionStr, "Version: 1.0.0") {
		t.Errorf("Version string does not contain correct version: %s", versionStr)
	}
	if !strings.Contains(versionStr, "Commit: abc123") {
		t.Errorf("Version string does not contain correct commit: %s", versionStr)
	}
	if !strings.Contains(versionStr, "Build Date: 2023-07-01") {
		t.Errorf("Version string does not contain correct build date: %s", versionStr)
	}

	// Restore original values
	Version = origVersion
	Commit = origCommit
	BuildDate = origBuildDate
}

func TestGetShortVersion(t *testing.T) {
	// Save original value
	origVersion := Version

	// Set test value
	Version = "1.0.0"

	// Check short version
	if v := GetShortVersion(); v != "1.0.0" {
		t.Errorf("Expected short version to be 1.0.0, got %s", v)
	}

	// Restore original value
	Version = origVersion
}
