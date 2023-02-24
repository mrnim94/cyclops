package main

import (
	"cyclops/log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func init() {
	os.Setenv("APP_NAME", "cyclops")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

}

var watcher *fsnotify.Watcher

// main
func main() {

	// Checking that an environment variable is present or not.
	cyclops, ok := os.LookupEnv("LOOK_PATH")
	if !ok {
		cyclops = "/Users/mrnim/GolandProjects/cyclops/nim"
	} else {
		log.Info("Cyclops look at path ", cyclops)
	}
	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for
	// directories
	if err := filepath.Walk(cyclops, watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				log.Info("EVENT! ", event)

				// watch for errors
			case err := <-watcher.Errors:
				log.Error("ERROR", err)
			}
		}
	}()

	<-done
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {
	if err != nil {
		log.Error(err)
		return nil
	}

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	} else {
		log.Error("")
	}

	return nil
}
