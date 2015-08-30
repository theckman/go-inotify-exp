package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/inotify"
)

func main() {
	watcher, err := inotify.NewWatcher()

	if err != nil {
		log.Fatalf("watcher error: %s", err)
	}

	err = watcher.Watch("/tmp")

	if err != nil {
		log.Fatalf("watching error: %s", err)
	}

	for {
		select {
		case event := <-watcher.Event:
			fmt.Printf("event: %s\n", event)
		case err = <-watcher.Error:
			fmt.Printf("watch error: %s", err)
		}
	}
}
