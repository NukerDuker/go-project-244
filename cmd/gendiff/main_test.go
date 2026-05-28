package main

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCLI_HelpFlagTest(t *testing.T) {
	tests := []struct {
		name     string
		argument string
	}{
		{name: "short help flag", argument: "-h"},
		{name: "long help flag", argument: "--help"},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		err := RunCLI(context.Background(), []string{"gendiff", tt.argument}, &buf)
		fmt.Println(buf.String())
		require.NoError(t, err, "Unexpected error in test")
		require.Contains(
			t,
			buf.String(),
			"--help, -h                  show help",
			"Output should contains help flag msg")
		require.Contains(
			t,
			buf.String(),
			"--format string, -f string  output format (default: \"stylish\")",
			"Output should contains format flag msg")
	}
}
