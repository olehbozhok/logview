package main

import (
	"fmt"
	"strings"

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
	path = strings.TrimPrefix(path, prefix)
	return fs.Box.Has(path)
}
