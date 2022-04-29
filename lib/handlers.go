package lib

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type APIJSONPictureDeleteRequest struct {
	PictureFilename string `json:"picture_filename"`
}

type APIJSONBasicResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type PageData struct {
	Images []string
}

func SetupHTTPHandlers() {

	picserve := http.FileServer(http.Dir(path))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", picserve))

	assetserve := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", assetserve))

	http.HandleFunc("/api/picture/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			{
				var jd APIJSONPictureDeleteRequest
				if r.Body == nil {
					http.Error(w, "No request body.", http.StatusBadRequest)
					return
				}
				err := json.NewDecoder(r.Body).Decode(&jd)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				if jd.PictureFilename == "" {
					http.Error(w, "Request did not include valid picture_filename.", http.StatusBadRequest)
					return
				}

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
	})

	indexTmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("template error:", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{Images: images}
		indexTmpl.Execute(w, data)
	})
}
