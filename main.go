package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// Generate AWS API mocks running go generate
//go:generate mockgen -source ./vendor/github.com/aws/aws-sdk-go/service/ecs/ecsiface/interface.go -package sdk -destination ./mock/aws/sdk/ecsiface_mock.go
//go:generate mockgen -source ./vendor/github.com/aws/aws-sdk-go/service/ec2/ec2iface/interface.go -package sdk -destination ./mock/aws/sdk/ec2iface_mock.go

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
		logrus.Errorf("Error creating garbage colletor: %s", err)
		return 1
	}

	logrus.Infof("Ready to rock")

	// Start the garbage collector if wanted
	if !cfg.disableGC {
		go gc.Run()
	} else {
		logrus.Warningf("Garbage collector is disabled, not running it!")
	}

	// Start the watcher loop
	err = w.Run()
	if err != nil {
		logrus.Errorf("Error after starting watcher: %s", err)
		return 1
	}

	return 0
}
