package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPath_positive(t *testing.T) {
	expRes := "{\"  host\":\"hexlet.io\",\"+ timeout\":20,\"+ verbose\":true,\"- follow\":false,\"- proxy\":\"123.234.53.22\",\"- timeout\":50}"
	result, err := GenDiff("testdata/file1.json", "testdata/file2.json", "stylish")
	require.NoError(t, err, "Test should be positive")
	require.Equal(t, expRes, result)
}
