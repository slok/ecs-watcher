package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

func main() {
	os.Exit(Main())
}

// Main will run the main program
func Main() int {
	// Parse command line flags
	if err := parse(os.Args[1:]); err != nil {
		logrus.Error(err)
		return 1
	}
	cfg := gCfg

	if cfg.debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Create the watcher
	w, err := NewWatcher(cfg)
	if err != nil {
		logrus.Errorf("Error creating watcher: %s", err)
		return 1
	}

	gc, err := NewGC(cfg)
	if err != nil {
		logrus.Errorf("Error creating garbage colletion: %s", err)
		return 1
	}
	go gc.Run()

	logrus.Infof("Ready to rock")

	// Start the garbage collector

	// Start the watcher loop
	err = w.Run()
	if err != nil {
		logrus.Errorf("Error after starting watcher: %s", err)
		return 1
	}

	return 0
}
