package lib

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var extensions = [...]string{".jpg", ".gif", ".png", ".webp"}

type Picture struct {
	Filename  string
	FullPath  string
	Directory *Directory
	Hash      string
}

type TPictures map[string]Picture

var Pictures = make(TPictures)

func FindPictures() error {

	for _, dir := range Directories {
		err := filepath.Walk(dir.Path, walker)
		if err != nil {
			return err
		}
	}

	clearEmptyDirectories()

	return nil
}

func clearEmptyDirectories() {
	hasPictures := make(map[string]bool)
	for _, pic := range Pictures {
		_, ok := hasPictures[pic.Directory.Path]
		if !ok {
			hasPictures[pic.Directory.Path] = true
		}
	}

	for k := range Directories {
		_, ok := hasPictures[k]
		if !ok {
			delete(Directories, k)
		}
	}
}

func walker(path string, info fs.FileInfo, err error) error {
	if isImage(path) {
		// Lookup the hash
		d := Directories[filepath.Dir(path)]
		Pictures[path] = Picture{
			Filename:  info.Name(),
			FullPath:  path,
			Directory: &d,
			Hash:      HashString(path),
		}
	} else if IsValidDir(path) {
		if IsDirectoryUnique(path) {
			Directories[path] = NewDirectory(path)
		}
	}
	return nil
}

// Check whether extensions exists in the filename
func isImage(filename string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}
