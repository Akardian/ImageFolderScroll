package image

import (
	"container/list"
	"context"
	"fmt"
	"image-scroll/config"
	"io/fs"
	"time"
)

// Check for new files in folder until connection is done
func RunFolderCheck(ctx context.Context) {
	imageList := list.New()

	// Init time with the last found file
	file := GetLastFileInfo(config.PATH)
	oldTime := GetLastImageModTime(file)

	fmt.Printf("Start With Image: [Name: %v, Time: %v]\n", file.Name(), file.ModTime())

	// Set Ticker to check for new file every second
	tickerFolderCheck := time.NewTicker(config.TICKER_FOLDER_CHECK)
	tickerSendNewImage := time.NewTicker(config.TICKER_SEND_NEW_IMAGE)

	for {
		select {
		case <-ctx.Done():
			tickerFolderCheck.Stop()
			close(SSEImageChannel)
			return
		case <-tickerFolderCheck.C:
			oldTime = addNewFiles(imageList, oldTime)
		case <-tickerSendNewImage.C:
			if imageList.Len() > 0 {
				sendEvent(imageList)
			}
		}
	}
}

// Check for new files and add them to the list
func addNewFiles(list *list.List, oldTime time.Time) time.Time {
	newFiles := CheckForNewFiles(config.PATH, oldTime)
	if len(newFiles) <= 0 {
		return oldTime
	}

	for _, file := range newFiles {
		fmt.Printf("NewFile [Name: %v, Time: %v]\n", file.Name(), file.ModTime())
		list.PushBack(file)
	}

	return newFiles[len(newFiles)-1].ModTime()
}

// Remove the oldest entry form the list and send a event to SSEImageChannel
func sendEvent(list *list.List) {
	eventData := list.Remove(list.Front()).(fs.FileInfo)

	if eventData != nil {
		fmt.Println("Send new image " + eventData.Name() + " to SSE channel")
		SSEImageChannel <- eventData.Name()
	}
}
