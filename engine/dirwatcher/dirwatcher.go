package dirwatcher

import (
	"github.com/didip/didip.github.io/libstring"
	"github.com/go-fsnotify/fsnotify"
	"log"
)

func WatchDir(path string, callback func(fsnotify.Event)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					callback(event)
				}

			case err := <-watcher.Errors:
				if err != nil {
					log.Println("Error: ", err)
				}
			}
		}
	}()

	err = watcher.Add(libstring.ExpandTildeAndEnv(path))
	if err != nil {
		return err
	}

	<-done
	return nil
}
