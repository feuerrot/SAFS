package safs

import (
	"net/http"

	"github.com/jessevdk/go-assets"
)

// StaticAssetsFS is an asset FS with some functions to use it for gin-contrib/staticgi
type StaticAssetsFS struct {
	FS *assets.FileSystem
}

// Exists is FIXME
func (f StaticAssetsFS) Exists(prefix string, path string) bool {
	if prefix != "/" {
		panic("We don't support prefixes except for the empty one")
	}

	_, ok := f.FS.Files[path]
	return ok
}

// Open is FIXME
func (f StaticAssetsFS) Open(name string) (http.File, error) {
	file, err := f.FS.Open(name)
	return file, err
}
