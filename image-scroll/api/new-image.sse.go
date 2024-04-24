package api

import (
	"container/list"
	"context"
	"fmt"
	"image-scroll/image"
	"image-scroll/sse"
	"io/fs"
	"net/http"
	"time"
)

var clientNumber = 0

func NewImageHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	connected()
	w.Header().Set("Content-Type", "text/event-stream")

	channel := make(chan string)
	go runCheck(r.Context(), channel)

	for imageName := range channel {
		event, err := sse.FormatServerSentEvent("new-image", imageName)
		if err != nil {
			fmt.Println(err)
			break
		}

		_, err = fmt.Fprint(w, event)
		if err != nil {
			fmt.Println(err)
			break
		}

		flusher.Flush()
	}

	disconnected()
}

// Check for new files in folder until connection is done
func runCheck(ctx context.Context, channel chan<- string) {
	var list = list.New()

	ticker := time.NewTicker(time.Second)
	lastCheck := image.GetLastFile("./api").ModTime()

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			close(channel)
			return
		case <-ticker.C:
			lastCheck = addNewFiles(list, lastCheck)

			eventData := sendEvent(list)
			if eventData != nil {
				channel <- eventData.Name()
			}
		}
	}
}

// Remove the oldest entry form the list and return the fileInfo
func sendEvent(list *list.List) fs.FileInfo {
	if list.Len() <= 0 {
		return nil
	}

	fileInfo := list.Remove(list.Front())

	// TODO: Check if file ist image

	return fileInfo.(fs.FileInfo)
}

// Check for new files and add them to the list
func addNewFiles(list *list.List, lastCheck time.Time) time.Time {
	newFiles := image.CheckForNewFiles(".", lastCheck)

	for _, file := range newFiles {
		list.PushBack(file)
		fmt.Printf("NewFile [Name: %v, Time: %v]\n", file.Name(), file.ModTime())
	}

	if len(newFiles) > 0 {
		return newFiles[len(newFiles)-1].ModTime()
	}
	return lastCheck
}

// Increase connected clients and have a nice console print
func connected() {
	clientNumber++
	fmt.Printf("Client connected. Total Clients: %v \n", clientNumber)
}

// Decrease connected clients and have a nice console print
func disconnected() {
	clientNumber--
	fmt.Printf("Client disconnected. Total Clients: %v \n", clientNumber)
}
