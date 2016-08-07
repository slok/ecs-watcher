package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

// Watcher  represents a watcher that will run every X time checking the status of the
// ECS cluster
type Watcher struct {
	// The name of the cluster
	clusterName string

	// The interval time to check
	interval time.Duration

	// The checker of the cluster
	checker Checker
}

// NewWatcher creates anew watcher
func NewWatcher(cfg Config) (*Watcher, error) {
	w := &Watcher{
		clusterName: cfg.clusterName,
		interval:    cfg.checkInterval,
	}

	// TODO: Checker selection
	c, err := NewAgentChecker(w.clusterName, cfg.awsRegion, cfg.unhealthyTag, cfg.markAfter)
	if err != nil {
		return nil, err
	}

	w.checker = c
	return w, nil
}

// Run will run the watcher
func (w *Watcher) Run() error {
	if w.checker == nil {
		return fmt.Errorf("No checker active on the watcher")
	}

	logrus.Infof("Starting to watch '%s' cluster every %s", w.clusterName, w.interval)
	t := time.NewTicker(w.interval)

	for range t.C {
		if err := w.checker.Check(); err != nil {
			logrus.Errorf("Error checking instances: %s", err)
			continue
		}
		if err := w.checker.Mark(); err != nil {
			logrus.Errorf("Error marking instances: %s", err)
			continue
		}
	}
	return nil

}
