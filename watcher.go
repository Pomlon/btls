package main

import (
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

type WatchPath struct {
	PauseWatch bool
	Watcher    *watcher.Watcher
}

//NewPath returns a new WatchPath, already set up with ignores and a recursive path to watch
func NewPath(path string, ignores ...string) WatchPath {
	w := WatchPath{
		PauseWatch: false,
		Watcher:    watcher.New(),
	}

	//w.Watcher.SetMaxEvents(1)

	if err := w.Watcher.AddRecursive(path); err != nil {
		log.Fatalln(err)
	}

	for _, path := range ignores {
		w.Watcher.Ignore(path)
	}

	return w
}

func (w *WatchPath) Start() {
	go w.Watcher.Start(time.Millisecond * 100)
	//go w.EventHandler()
}

func (w *WatchPath) EventHandler() {
	for {
		select {
		case event := <-w.Watcher.Event:
			fmt.Println(event) // Print the event's info.
		case err := <-w.Watcher.Error:
			log.Fatalln(err)
		case <-w.Watcher.Closed:
			return
		}
	}
}

func (w *WatchPath) Pause() {
	w.PauseWatch = true
}

func (w *WatchPath) Unpause() {
	w.PauseWatch = false
}
