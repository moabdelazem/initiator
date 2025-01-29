package projects

import (
	"os"
	"testing"
)

func TestChangeDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Change to the temporary directory
	err := ChangeDirectory(tempDir)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if currentDir != tempDir {
		t.Fatalf("expected directory %s, got %s", tempDir, currentDir)
	}
}

func TestNewProject(t *testing.T) {
	projectName := "test_project"
	projectDir := t.TempDir()

	nodeProject := NewProject(projectName, projectDir, NodeJS)
	if _, ok := nodeProject.(*NodeProject); !ok {
		t.Fatalf("expected NodeProject, got %T", nodeProject)
	}

	goProject := NewProject(projectName, projectDir, GoLang)
	if _, ok := goProject.(*GoProject); !ok {
		t.Fatalf("expected GoProject, got %T", goProject)
	}

	unknownProject := NewProject(projectName, projectDir, "unknown")
	if unknownProject != nil {
		t.Fatalf("expected nil, got %T", unknownProject)
	}
}
