package bucketfs

import (
	"io/fs"
	"sync"
)

var (
	adapters     = make(map[string]Adapter)
	adaptersLock = &sync.RWMutex{}
)

// Adapter is the interface that must be implemented by an adapter.
type Adapter interface {
	// Open opens a bucketfs with the given name and options.
	Open(name string, opts map[string]string) (f fs.FS, err error)
}

// Open opens a bucketfs with the given name and options.
func Open(name string, opts map[string]string) (fs.FS, error) {
	adaptersLock.RLock()
	defer adaptersLock.RUnlock()
	if opts == nil {
		opts = make(map[string]string)
	}
	if adapter, ok := adapters[name]; ok {
		return adapter.Open(name, opts)
	} else {
		return nil, &AdapterNotFound{AdapterName: name}
	}
}

// Register registers an adapter with the given name.
func Register(name string, adapter Adapter) {
	adaptersLock.Lock()
	defer adaptersLock.Unlock()
	adapters[name] = adapter
}
