package pkg

import (
	"bytes"
	"os/exec"
	"strings"
)

// IsGoInstalled checks if Go is installed on the system
func IsGoInstalled() bool {
	_, err := exec.LookPath("go")
	return err == nil
}

// GetGoVersion returns the installed version of Go
func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Parse the output which is typically like "go version go1.16.5 darwin/amd64"
	version := strings.TrimSpace(out.String())
	parts := strings.Split(version, " ")
	if len(parts) >= 3 {
		return parts[2], nil
	}
	return version, nil
}

// IsNodeInstalled checks if Node.js is installed on the system
func IsNodeInstalled() bool {
	_, err := exec.LookPath("node")
	return err == nil
}

// GetNodeVersion returns the installed version of Node.js
func GetNodeVersion() (string, error) {
	cmd := exec.Command("node", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Clean up the output (typically like "v14.17.0")
	version := strings.TrimSpace(out.String())
	// Remove 'v' prefix if present
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}
	return version, nil
}

// IsGitInstalled checks if Git is installed on the system
func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// GetGitVersion returns the installed version of Git
func GetGitVersion() (string, error) {
	cmd := exec.Command("git", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Output is typically "git version 2.30.1"
	version := strings.TrimSpace(out.String())
	parts := strings.Split(version, " ")
	if len(parts) >= 3 {
		return parts[2], nil
	}
	return version, nil
}

// IsDockerInstalled checks if Docker is installed on the system
func IsDockerInstalled() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}

// GetDockerVersion returns the installed version of Docker
func GetDockerVersion() (string, error) {
	cmd := exec.Command("docker", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Output is typically "Docker version 20.10.7, build f0df350"
	version := strings.TrimSpace(out.String())
	parts := strings.Split(version, " ")
	if len(parts) >= 3 {
		// Remove trailing comma if present
		return strings.TrimSuffix(parts[2], ","), nil
	}
	return version, nil
}

// IsDockerComposeInstalled checks if Docker Compose is installed on the system
func IsDockerComposeInstalled() bool {
	// First try docker-compose command (older versions)
	_, err := exec.LookPath("docker-compose")
	if err == nil {
		return true
	}

	// Then try docker compose command (newer versions)
	// Check if docker is installed first (required for 'docker compose')
	_, err = exec.LookPath("docker")
	if err != nil {
		return false
	}

	// Check if 'docker compose' works
	cmd := exec.Command("docker", "compose", "version")
	err = cmd.Run()
	return err == nil
}

// GetDockerComposeVersion returns the installed version of Docker Compose
func GetDockerComposeVersion() (string, error) {
	var cmd *exec.Cmd

	// Check if standalone docker-compose is available
	_, err := exec.LookPath("docker-compose")
	if err == nil {
		cmd = exec.Command("docker-compose", "version", "--short")
	} else {
		// Try docker compose (newer version)
		cmd = exec.Command("docker", "compose", "version", "--short")
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	version := strings.TrimSpace(out.String())
	// Remove 'v' prefix if present
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}
	return version, nil
}
