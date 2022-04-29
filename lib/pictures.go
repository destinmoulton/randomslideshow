package lib

import (
	"path/filepath"
	"strings"
)

var extensions = [...]string{".jpg", ".gif", ".png", ".webp"}

type Picture struct {
	Filename string
	FullPath string
}

func getPictures(path string) ([]string, error) {

	results, err := filepath.Glob(path)

	if err != nil {
		return nil, err
	}
	var pics []string
	for _, path := range results {
		if isImage(path) {
			_, file := filepath.Split(path)
			pics = append(pics, file)
		}
	}
	return pics, nil
}

/*
func walker(path string, info fs.FileInfo, err error) error {

}
*/

// Check whether extensions exists in the filename
func isImage(filename string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}
