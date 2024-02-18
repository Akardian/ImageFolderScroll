package api

import (
	"log"
	"net/http"
	"os"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {

	buf, err := os.ReadFile("api/sid.jpeg")

	if err != nil {

		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/*")
	w.Header().Set("Content-Disposition", `attachment;filename="sid.jpeg"`)

	w.Write(buf)
}
