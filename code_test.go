package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPath_positive(t *testing.T) {
	expRes := "{\n  - follow: false\n    host: hexlet.io\n  - proxy: 123.234.53.22\n  - timeout: 50\n  + timeout: 20\n  + verbose: true\n}"
	result, err := GenDiff("testdata/file1.json", "testdata/file2.json", "stylish")
	require.NoError(t, err, "Test should be positive")
	require.Equal(t, expRes, result)
}
