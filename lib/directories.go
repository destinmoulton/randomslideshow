package lib

import (
	"os"
)

type Directory struct {
	Hash string
	Path string
}

var Directories = make(map[string]Directory)

func NewDirectory(path string) Directory {
	hash := HashString(path)
	return Directory{hash, path}
}

// Does this directory not exist in the map?
func IsDirectoryUnique(dir string) bool {
	_, ok := Directories[dir]
	return !ok
}

// Is this path a directory?
func IsValidDir(dir string) bool {

	dirinfo, err := os.Stat(dir)

	return !os.IsNotExist(err) && dirinfo.IsDir()
}
