package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"randomslideshow/lib"
)

var extensions = [...]string{".jpg", ".gif", ".png", ".webp"}

func main() {

	if len(os.Args) <= 1 {
		log.Panic("You must specify a directory where images are stored.")
	}

	path := os.Args[1]

	dirinfo, err := os.Stat(path)

	if os.IsNotExist(err) || !dirinfo.IsDir() {
		log.Panic("That directory does not exist or has an error.")
	}

	fullglob := filepath.Join(path, "*")
	fmt.Printf("Loading glob %s", fullglob)

	results, err := filepath.Glob(fullglob)

	if err != nil {
		log.Panic("Glob failed to search path.", err)
	}

	var images []string
	for _, path := range results {
		if isImage(path) {
			_, file := filepath.Split(path)
			images = append(images, file)
		}
	}
	lib.SetupHTTPHandlers()
	http.ListenAndServe(":3050", nil)

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
