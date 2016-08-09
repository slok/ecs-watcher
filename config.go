package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"
)

// Default configuration
const (
	defaultDebug         = false
	defaultCheckInterval = 5 * time.Second
	defaultGCInterval    = 2 * time.Second
	defaultMarkAfter     = 1 * time.Minute
	defaultStepPercent   = 20
	defaultUnhealthyTag  = "unhealthy:true"
	defaultDisableGC     = false
)

// Config represents the main configuration
type Config struct {
	fs *flag.FlagSet

	clusterName   string
	debug         bool
	checkInterval time.Duration
	gcInterval    time.Duration
	gcStepPercent int
	awsRegion     string
	unhealthyTag  string
	markAfter     time.Duration
	disableGC     bool
}

var gCfg = Config{}

// init will load all the cmd flags
func init() {
	gCfg.fs = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	gCfg.fs.StringVar(
		&gCfg.clusterName, "cluster", "",
		"The target cluster name",
	)

	gCfg.fs.StringVar(
		&gCfg.awsRegion, "region", "",
		"The AWS region of the cluster",
	)

	gCfg.fs.DurationVar(
		&gCfg.checkInterval, "check.interval", defaultCheckInterval,
		"The interval for checking the cluster",
	)

	gCfg.fs.DurationVar(
		&gCfg.gcInterval, "gc.interval", defaultGCInterval,
		"The minimum interval for garbage collection of unhealthy targets",
	)

	gCfg.fs.DurationVar(
		&gCfg.markAfter, "unhealthy.after", defaultMarkAfter,
		"The duration that a target needs to be unhealthy to declare as unhealthy",
	)

	gCfg.fs.IntVar(
		&gCfg.gcStepPercent, "gc.step.percent", defaultStepPercent,
		"The step percent of total unhealthy targets when cleaning",
	)

	gCfg.fs.StringVar(
		&gCfg.unhealthyTag, "unhealthy.tag", defaultUnhealthyTag,
		"The tag used to mark unhealty labels key:value form",
	)

	gCfg.fs.BoolVar(
		&gCfg.debug, "debug", defaultDebug,
		"Run in debug mode",
	)

	gCfg.fs.BoolVar(
		&gCfg.disableGC, "disable.gc", defaultDisableGC,
		"Don't run garbage collector",
	)
}

func parse(args []string) error {

	if err := gCfg.fs.Parse(args); err != nil {
		return err
	}
	match, err := regexp.MatchString(`^[^:]+:[^:]+$`, gCfg.unhealthyTag)
	if !match || err != nil {
		return fmt.Errorf("Wrong tag format, must be key:value format. Help: %s -h", os.Args[0])
	}

	if gCfg.awsRegion == "" {
		return fmt.Errorf("Cluster AWS region must be set. Help: %s -h", os.Args[0])
	}

	if gCfg.clusterName == "" {
		return fmt.Errorf("Cluster name can't be empty. Help: %s -h", os.Args[0])
	}
	if len(gCfg.fs.Args()) != 0 {
		return fmt.Errorf("Invalid command line arguments. Help: %s -h", os.Args[0])

	}
	return nil
}
