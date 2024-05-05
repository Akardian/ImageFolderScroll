package api

import (
	"image-scroll/config"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

func GetImageHandler(w http.ResponseWriter, r *http.Request) {

	page, _ := template.New("Template").Parse("<img src='api/image'</img>")
	page.Execute(w, nil)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	name := config.PATH + "/" + path.Base(r.RequestURI)
	buf, err := os.ReadFile(name)

	if err != nil {
		log.Fatal("Image Handler: ", err)
	}

	w.Header().Set("Content-Type", "image/*")
	w.Header().Set("Content-Disposition", `attachment;filename="sid.jpeg"`)

	w.Write(buf)
}
