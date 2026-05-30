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

func TestArgumentsCount_positive(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "one argument passed", args: []string{"gendiff", "some/path/1"}},
		{name: "zero argument passed", args: []string{"gendiff"}},
	}
	var buf bytes.Buffer
	for _, tt := range tests {
		actualErr := RunCLI(context.Background(), tt.args, &buf)
		expErr := fmt.Errorf("gendiff needs two arguments to compare files, actual arguments number: %d",
			len(tt.args)-1)
		require.Errorf(t, actualErr, "test should produce an error")
		require.Equal(t, expErr, actualErr, "error should be equal to expected error")
	}
}
