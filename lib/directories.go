package lib

import (
	"crypto/sha1"
	"fmt"
	"os"
)

type Directory struct {
	Hash string
	Path string
}

var Directories = make(map[string]Directory)

func NewDirectory(path string) Directory {
	hash := HashDirectory(path)
	return Directory{hash, path}
}

// Get the sha1 hash hex of a directory
func HashDirectory(dir string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(dir)))
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
