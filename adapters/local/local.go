package local

import (
	"io/fs"
	"os"

	"github.com/yankeguo/bucketfs"
)

type localAdapter struct{}

func (a *localAdapter) Open(name string, opts map[string]string) (fs.FS, error) {
	path := opts["path"]
	if path == "" {
		return nil, &bucketfs.AdapterOptionMissing{OptionName: name}
	}
	return os.DirFS(path), nil
}

func init() {
	bucketfs.Register("local", &localAdapter{})
}
