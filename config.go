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
	defaultDebug        = false
	defaultInterval     = 5 * time.Second
	defaultUnhealthyTag = "unhealthy:true"
)

var cfg = struct {
	fs *flag.FlagSet

	clusterName  string
	debug        bool
	interval     time.Duration
	awsRegion    string
	unhealthyTag string
}{}

// init will load all the cmd flags
func init() {
	cfg.fs = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	cfg.fs.StringVar(
		&cfg.clusterName, "cluster", "",
		"The target cluster name",
	)

	cfg.fs.StringVar(
		&cfg.awsRegion, "region", "",
		"The AWS region of the cluster",
	)

	cfg.fs.DurationVar(
		&cfg.interval, "interval", defaultInterval,
		"The interval of checking the cluster",
	)
	cfg.fs.StringVar(
		&cfg.unhealthyTag, "unhealthy.tag", defaultUnhealthyTag,
		"The tag used to mark unhealty labels key:value form",
	)

	cfg.fs.BoolVar(
		&cfg.debug, "debug", defaultDebug,
		"Run in debug mode",
	)
}

func parse(args []string) error {

	if err := cfg.fs.Parse(args); err != nil {
		return err
	}
	match, err := regexp.MatchString(`^[^:]+:[^:]+$`, cfg.unhealthyTag)
	if !match || err != nil {
		return fmt.Errorf("Wrong tag format, must be key:value format. Help: %s -h", os.Args[0])
	}

	if cfg.awsRegion == "" {
		return fmt.Errorf("Cluster AWS region must be set. Help: %s -h", os.Args[0])
	}

	if cfg.clusterName == "" {
		return fmt.Errorf("Cluster name can't be empty. Help: %s -h", os.Args[0])
	}

	if len(cfg.fs.Args()) != 0 {
		return fmt.Errorf("Invalid command line arguments. Help: %s -h", os.Args[0])

	}
	return nil
}