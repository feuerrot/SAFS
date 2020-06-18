package main

import (
	"net/http"

	"github.com/jessevdk/go-assets"
)

// StaticAssetsFS is FIXME
type StaticAssetsFS struct {
	fs *assets.FileSystem
}

// Exists is FIXME
func (f StaticAssetsFS) Exists(prefix string, path string) bool {
	if prefix != "/" {
		panic("We don't support prefixes except for the empty one")
	}

	_, ok := f.fs.Files[path]
	return ok
}

// Open is FIXME
func (f StaticAssetsFS) Open(name string) (http.File, error) {
	file, err := f.fs.Open(name)
	return file, err
}
