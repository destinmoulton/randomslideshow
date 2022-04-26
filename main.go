package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

var extensions = [...]string{".jpg", ".gif", ".png", ".webp"}

func main() {
	picPath := ""
	glb := "{*.jpg,*.png,*.gif,*.webp}"
	if len(os.Args) > 1 {
		picPath = os.Args[1]
	} else {
		picPath = "./"
	}
	picPath, err := filepath.Abs(picPath)
	if err != nil {
		log.Panic("Failed to get absolute path", err)
	}
	fmt.Printf("Loading pictures from %s", picPath)

	fullglob := filepath.Join(picPath, glb)
	fmt.Printf("Loading glob %s", fullglob)

	results, err := filepath.Glob(fullglob)

	if err != nil {
		log.Panic("Glob failed to search path.", err)
	}

	for _, file := range results {
		fmt.Println(file)
	}

	myApp := app.New()
	w := myApp.NewWindow("Image")

	//image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	//image := canvas.NewImageFromImage("ramsey.jpg")
	// image := canvas.NewImageFromReader(reader, name)
	image := canvas.NewImageFromFile("ramsey.jpg")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}

// Check whether extensions exists in the filename
func isImage(filename string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(ext) {
			return true
		}
	}
	return false
}
