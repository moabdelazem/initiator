package projects

import (
	"os"
	"os/exec"
	"testing"
)

var execCommand = exec.Command

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	os.Exit(0)
}
