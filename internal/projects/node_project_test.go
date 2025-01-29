package projects

import (
	"os"
	"os/exec"
	"testing"
)

func TestNodeProject_Create(t *testing.T) {
	projectName := "test_node_project"
	projectDir := t.TempDir()

	nodeProject := &NodeProject{Name: projectName, Dir: projectDir}

	// Mock the exec.Command function
	execCommand = func(name string, arg ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcess", "--", name}
		cs = append(cs, arg...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}

	err := nodeProject.Create()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
