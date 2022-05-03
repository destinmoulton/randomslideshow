package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type APIJSONPictureDeleteRequest struct {
	Action      string `json:"action"`
	PicturePath string `json:"picture_path"`
}

type APIJSONBasicResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type PageData struct {
	Images TPictures
}

func SetupHTTPHandlers() {

	for _, dir := range Directories {
		picserve := http.FileServer(http.Dir(dir.Path))
		b := fmt.Sprintf("/%s/", dir.Hash)
		http.Handle(b, http.StripPrefix(b, picserve))
	}

	assetserve := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", assetserve))

	http.HandleFunc("/api/picture/", apiPictureHandler)

	indexTmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("template error:", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{Images: Pictures}
		indexTmpl.Execute(w, data)
	})
}

func apiPictureHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodDelete:
		{
			var jd APIJSONPictureDeleteRequest
			if r.Body == nil {
				http.Error(w, "No request body.", http.StatusBadRequest)
				return
			}
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields()
			err := decoder.Decode(&jd)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if jd.PicturePath == "" {
				http.Error(w, "Request did not include valid picture_filename.", http.StatusBadRequest)
				return
			}
			if _, err := os.Stat(jd.PicturePath); errors.Is(err, os.ErrNotExist) {
				http.Error(w, "That file does not exist.", http.StatusBadRequest)
				return
			}

			err = os.Remove(jd.PicturePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			RemovePictureFromMap(jd.PicturePath)

			data := APIJSONBasicResponse{200, "Successfully deleted picture."}
			msg, err := json.Marshal(data)
			if err != nil {
				// TODO: Change the error message to something unrevealing
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(msg)
		}
	}
}
