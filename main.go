package main

import (
	"fmt"
	"net/http"
	"os"

	"randomslideshow/lib"
)

func main() {

	flags, err := lib.ParseCLIArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = lib.FindPictures()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	lib.SetupHTTPHandlers()
	h := fmt.Sprintf("%s:%s", flags["ip"], flags["port"])
	fmt.Printf("Server listening on: http://%s\n", h)
	http.ListenAndServe(h, nil)
}
