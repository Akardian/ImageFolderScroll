package image

import (
	"fmt"
	"image-scroll/config"
	"image-scroll/sse"
	"net/http"
)

var ClientNumber = 0
var SSEImageChannel = make(chan string)

func SSEImageHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	connected(flusher, w)

	// Create and send SSE on Channel activity
	for {
		select {
		case <-r.Context().Done():
			disconnected()
			return
		case <-SSEImageChannel:
			sendSSE(flusher, w, <-SSEImageChannel)
		}
	}

}

// Increase connected clients and have a nice console print
func connected(flusher http.Flusher, w http.ResponseWriter) {
	ClientNumber++

	file := GetLastFileInfo(config.PATH)
	if file != nil {
		sendSSE(flusher, w, file.Name())
	}

	fmt.Printf("Client connected. Total Clients: %v \n", ClientNumber)
}

// Decrease connected clients and have a nice console print
func disconnected() {
	ClientNumber--
	fmt.Printf("Client disconnected. Total Clients: %v \n", ClientNumber)
}

func sendSSE(flusher http.Flusher, w http.ResponseWriter, event string) {
	event, err := sse.FormatServerSentEvent("new-image", event)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprint(w, event)
	if err != nil {
		fmt.Println(err)
		return
	}

	flusher.Flush()

}
