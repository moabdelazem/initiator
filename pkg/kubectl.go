package pkg

import (
	"os/exec"
	"strings"
)

// IsKubectlInstalled checks if kubectl is installed
func IsKubectlInstalled() bool {
	_, err := exec.LookPath("kubectl")
	return err == nil
}

// GetKubectlVersion returns the installed kubectl version
func GetKubectlVersion() (string, error) {
	cmd := exec.Command("kubectl", "version", "--client", "--short")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	// Clean up the output to extract just the version
	version := strings.TrimSpace(string(output))
	// Extract client version from output that looks like: "Client Version: v1.23.5"
	version = strings.TrimPrefix(version, "Client Version: ")
	return version, nil
}
