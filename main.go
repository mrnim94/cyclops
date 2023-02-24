package main

import (
	"cyclops/log"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	"sync"
)

func init() {
	os.Setenv("APP_NAME", "cyclops")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

}

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
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Watch the specified directory and all subdirectories.
	err = filepath.Walk(cyclops, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error(err)
			return err
		}
		var mutex = &sync.Mutex{}

		mutex.Lock()
		defer mutex.Unlock()

		if info.IsDir() {
			err := watcher.Add(path)
			if err != nil {
				log.Error(err)
			}
			log.Info("Found Directory:", path)
		} else {
			log.Info("Found file:", path)
		}

		return nil
	})
	if err != nil {
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
