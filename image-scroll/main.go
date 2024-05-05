package main

import (
	"context"
	"fmt"
	"image-scroll/api"
	"image-scroll/config"
	"image-scroll/image"
	"log"
	"net/http"
)

func main() {
	// Start folder check
	ctx := context.Background()
	go image.RunFolderCheck(ctx)

	// Start SSE Image handler
	http.HandleFunc("/image/update", image.SSEImageHandler)

	// Start HTML provider
	http.HandleFunc("/", api.PageHandler)
	http.HandleFunc("/htmx", api.HTMXHandler)

	http.HandleFunc("/image/{id}", api.ImageHandler)
	http.HandleFunc("/image/new", api.GetImageHandler)

	fmt.Println("Listening to port: " + config.PORT)
	log.Fatal(http.ListenAndServe(config.PORT, nil))
}
