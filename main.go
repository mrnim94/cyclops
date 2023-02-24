package main

import (
	"cyclops/log"
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
		log.Error("Error walking directory:", err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Info("event:", event)
				if event.Has(fsnotify.Write) {
					log.Info("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error("error:", err)
			}
		}
	}()

	// Block main goroutine forever.
	<-make(chan struct{})
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {
	if err != nil {
		// Ignore "no such file or directory" errors
		if os.IsNotExist(err) {
			log.Error("Ignored ", err)
			return nil
		}
		log.Error("Error accessing file or directory:", err)
		return err
	}

	if fi.IsDir() {
		err := watcher.Add(path)
		if err != nil {
			log.Error(err)
		}
		log.Info("Found Directory:", path)
	} //else {
	//	log.Info("Found file:", path)
	//}

	return nil
}
