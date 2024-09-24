package local

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yankeguo/bucketfs"
)

func TestLocalAdapter(t *testing.T) {
	// Register the adapter
	bucketfs.Register("local", &localAdapter{})

	// Open the adapter
	_, err := bucketfs.Open("local", map[string]string{"no_path": "."})
	require.Error(t, err)
	require.IsType(t, &bucketfs.AdapterOptionMissing{}, err)

	// Open the adapter
	fs, err := bucketfs.Open("local", map[string]string{"path": "."})
	require.NoError(t, err)

	// Read the file
	f, err := fs.Open("local_test.go")
	require.NoError(t, err)
	f.Close()
}
