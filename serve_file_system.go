package main

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gobuffalo/packr/v2"
)

var _ static.ServeFileSystem = &ServeFileSystem{}

// ServeFileSystem wraper around static.ServeFileSystem
type ServeFileSystem struct {
	*packr.Box
}

// Exists return true if file exist
func (fs *ServeFileSystem) Exists(prefix string, path string) bool {
	fmt.Println(prefix, path)
	return fs.Box.Has(path)
}
