package main

import (
	"cyclops/log"
	"github.com/fsnotify/fsnotify"
	"io/fs"
	"os"
	"strings"
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
		cyclops = "/Users/mrnim/GolandProjects/cyclops/nim,/Users/mrnim/GolandProjects/cyclops/nim/chirldren"
	} else {
		log.Info("Cyclops look at paths ", cyclops)
	}
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
	}
	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for
	// directories
	//if err := filepath.WalkDir(cyclops, watchDir); err != nil {
	//	log.Error("Error walking directory:", err)
	//}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Info("channel closed")
					return
				} else {
					log.Info("action:", event.Op.String(), " --> ", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Info("channel closed")
					return
				} else {
					log.Fatal("error:", err)
				}
			}
		}
	}()

	watchDirectories := strings.Split(cyclops, ",")
	for _, watchDirectory := range watchDirectories {
		// Add a path.
		err := watcher.Add(watchDirectory)
		if err != nil {
			log.Fatal("Add folder", watchDirectory, " to watch is error --> ", err)
		} else {
			log.Info("Cyclops look at path ", watchDirectory)
		}
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi fs.DirEntry, err error) error {
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
