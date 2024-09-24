package bucketfs

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type dummyAdapter struct {
	f   fs.FS
	err error
}

func (a *dummyAdapter) Open(name string, opts map[string]string) (f fs.FS, err error) {
	return a.f, a.err
}

func TestRegisterOpen(t *testing.T) {
	f := os.DirFS(".")
	err := errors.New("dummy error")
	Register("test", &dummyAdapter{
		f:   f,
		err: err,
	})
	f1, err2 := Open("test", nil)
	require.Equal(t, f, f1)
	require.Equal(t, err, err2)
}
