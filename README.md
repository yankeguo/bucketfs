# bucketfs

Adapter to `Golang` `fs.FS` for various object storage services

## Usage

1. Import a adapter

    ```go
    import (
        "github.com/yankeguo/bucketfs"
        _ "github.com/yankeguo/bucketfs/adapters/local"
    )
    ```

2. Open a bucketfs

    ```go
    func main() {
        fs, err := bucketfs.New("local", map[string]string{
            "path": "/path/to/root",
        })
        if err != nil {
            panic(err)
        }
        // use fs
    }
    ```

## Adapters

### `local`

Adapter to local file system

**Options**

* `path`, required, the root path of the local file system

## Credits

GUO YANKE, MIT License
