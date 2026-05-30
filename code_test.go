package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPath_positive(t *testing.T) {
	result, err := GenDiff("testdata/file1.json", "testdata/file2.json", "stylish")
	require.NoError(t, err, "Test should be positive")
	require.Contains(t, result, "host:hexlet.io")
}
