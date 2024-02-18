package image

import (
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

var folderRegex = regexp.MustCompile("[^/]+$")

// Check the given folder for new files and return an array of fs.FileInfo for all new files
func CheckForNewFiles(root string, lastTime time.Time) []fs.FileInfo {
	rootFolder := folderRegex.FindString(root)

	var newFiles []fs.FileInfo

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Could not find folder '" + rootFolder + "'")
		}

		if info.IsDir() && info.Name() != rootFolder {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		if info.ModTime().After(lastTime) {
			newFiles = append(newFiles, info)
		}

		return nil
	})

	sort.Slice(newFiles, func(a int, b int) bool {
		return newFiles[a].ModTime().Before(newFiles[b].ModTime())
	})
	return newFiles
}

// Returns the last modified file in the folder
func GetLastFile(root string) fs.FileInfo {
	rootFolder := folderRegex.FindString(root)

	var lastFile fs.FileInfo = nil

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Could not find folder '" + rootFolder + "'")
		}

		if info.IsDir() && info.Name() != rootFolder {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		if lastFile == nil || info.ModTime().After(lastFile.ModTime()) {
			lastFile = info
		}

		return nil
	})

	return lastFile
}
