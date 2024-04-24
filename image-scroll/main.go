package main

import (
	"fmt"
	"image-scroll/api"
	"log"
	"net/http"
)

var PORT = ":8000"

func main() {
	http.HandleFunc("/", api.PageHandler)
	http.HandleFunc("/htmx", api.HTMXHandler)

	http.HandleFunc("/image/update", api.NewImageHandler)

	http.HandleFunc("/image/{id}", api.ImageHandler)
	http.HandleFunc("/image/new", api.GetImageHandler)

	fmt.Println("Listening to port: " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
