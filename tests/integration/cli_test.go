package integration

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func getBinaryPath() string {
	wd, _ := os.Getwd()
	return filepath.Join(wd, "..", "..", "calculator")
}

// TestInteractiveMode_Entry verifies interactive mode starts correctly
func TestInteractiveMode_Entry(t *testing.T) {
	input := "10 / 2\nexit\n"
	inputFile, err := ioutil.TempFile("", "calc_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	if _, err := inputFile.WriteString(input); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	inputFile.Close()

	cmd := exec.Command(getBinaryPath())
	cmd.Stdin, _ = os.Open(inputFile.Name())

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("Process ended with: %v", err)
	}

	outputStr := string(output)
	t.Logf("Output: %q", outputStr)
	if !strings.Contains(outputStr, ">") {
		t.Error("Expected '>' prompt in output")
	}
	if !strings.Contains(outputStr, "= 5") {
		t.Errorf("Expected '= 5' in output, got '%s'", outputStr)
	}
}

// TestInteractiveMode_SequentialCalculations verifies multiple calculations work
func TestInteractiveMode_SequentialCalculations(t *testing.T) {
	input := "10 / 2\n7 - 3\n4 * 2\nexit\n"
	inputFile, err := ioutil.TempFile("", "calc_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	if _, err := inputFile.WriteString(input); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	inputFile.Close()

	cmd := exec.Command(getBinaryPath())
	cmd.Stdin, _ = os.Open(inputFile.Name())

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("Process ended with: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "= 5") {
		t.Errorf("Expected '= 5' in output")
	}
	if !strings.Contains(outputStr, "= 4") {
		t.Errorf("Expected '= 4' in output")
	}
	if !strings.Contains(outputStr, "= 8") {
		t.Errorf("Expected '= 8' in output")
	}
}

// TestInteractiveMode_EmptyLineHandling verifies empty lines are handled
func TestInteractiveMode_EmptyLineHandling(t *testing.T) {
	input := "\n10 / 2\nexit\n"
	inputFile, err := ioutil.TempFile("", "calc_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	if _, err := inputFile.WriteString(input); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	inputFile.Close()

	cmd := exec.Command(getBinaryPath())
	cmd.Stdin, _ = os.Open(inputFile.Name())

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("Process ended with: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "= 5") {
		t.Errorf("Expected '= 5' in output after empty line")
	}
}
