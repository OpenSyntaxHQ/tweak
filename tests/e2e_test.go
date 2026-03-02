package tests

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	binaryName := "tweak_e2e"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	buildCmd := exec.Command("go", "build", "-o", binaryName, "../main.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build binary for E2E tests: %v", err)
	}
	defer os.Remove(binaryName)

	binaryPath, err := filepath.Abs(binaryName)
	if err != nil {
		t.Fatalf("Failed to get absolute path for binary: %v", err)
	}

	tests := []struct {
		name           string
		args           []string
		stdin          string
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "Base64 Encoding from Args",
			args:           []string{"base64-encode", "hello world"},
			expectedOutput: "aGVsbG8gd29ybGQ=",
		},
		{
			name:           "Base64 Decoding from Stdin (Piping)",
			args:           []string{"base64-decode"},
			stdin:          "aGVsbG8gd29ybGQ=",
			expectedOutput: "hello world",
		},
		{
			name:           "MD5 Hash (Arg with spaces)",
			args:           []string{"md5", "tweak test"},
			expectedOutput: "4b284490e3437e4d4a48c09edb0040a0",
		},
		{
			name:           "Uppercase transform",
			args:           []string{"upper", "make me loud"},
			expectedOutput: "MAKE ME LOUD",
		},
		{
			name:        "Invalid Command Error",
			args:        []string{"fake-command-123"},
			expectError: true,
		},
		{
			name:        "AES Encrypt Missing Required Key",
			args:        []string{"aes-encrypt", "hello"},
			expectError: true,
		},
		{
			name:        "Regex Match Missing Required Pattern",
			args:        []string{"regex-match", "hello"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)

			// Set up stdin if provided (simulates bash/CMD piping: echo | tweak)
			if tt.stdin != "" {
				cmd.Stdin = strings.NewReader(tt.stdin)
			}

			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but command succeeded. output: %s", stdout.String())
				}
				return
			}

			if err != nil {
				t.Fatalf("Command failed: %v\nStderr: %s", err, stderr.String())
			}

			// Clean up output (CRLF to LF for cross-platform comparison)
			got := strings.TrimSpace(stdout.String())
			got = strings.ReplaceAll(got, "\r\n", "\n")

			if got != tt.expectedOutput {
				t.Errorf("\nGot:      %q\nExpected: %q", got, tt.expectedOutput)
			}
		})
	}
}
