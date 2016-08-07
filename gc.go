package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

// GC  represents the garbage collector of the unhleathy targets
type GC struct {

	// collection interval
	interval time.Duration

	// Cleaner
	cleaner Cleaner
}

// NewGC creates a new garbage collector
func NewGC(interval time.Duration) (*GC, error) {
	gc := &GC{
		interval: interval,
	}
	// TODO: clener selection
	k, err := NewKiller(cfg.awsRegion, cfg.gcStepPercent, cfg.unhealthyTag)
	if err != nil {
		return nil, err
	}

	gc.cleaner = k

	return gc, nil
}

// Run will start the garbage collector
func (g *GC) Run(dryRun bool) error {

	if g.cleaner == nil {
		return fmt.Errorf("No cleaner active on the garbage collector")
	}

	logrus.Infof("Starting garbage collector")
	t := time.NewTicker(g.interval)

	for range t.C {
		err := g.cleaner.Clean()
		if err != nil {
			logrus.Errorf("Error cleaning instances: %s", err)
			continue
		}
	}

	return nil
}
