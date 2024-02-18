package main

import (
	"image-scroll/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.PageHandler)
	http.HandleFunc("/crypto-price", api.NewImageHandler)
	http.HandleFunc("/image", api.ImageHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
