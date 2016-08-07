# ECS watcher

[![Build Status](https://travis-ci.org/slok/ecs-watcher.svg?branch=master)](https://travis-ci.org/slok/ecs-watcher)

ECS agent watcher will destroy the instances that are unhealthy. To know if
an instance is unhealthy on the cluster, ECS watcher will check periodically
the status of the instance ECS agent, if this agent is not connected for a given
time, this instance will be marked as unhealthy and the garbage collector will
kill the instance.

## Requirements:

+ Your instances should be in an autoscalation group, ECS watcher only kills instances, it doesn't run again
* Your automated instnaces should connect automatically to the ECS cluster
* YOur serices should be HA, this means that need to have more than one instances per service, remember that
ECS watcher will kill the unhealthy instances without notification

## Minimum usage example:

```bash
ecs-watcher --cluster="slok-ECSCluster1-15OBYPKBNXIO6" --region=us-west-2
```

## Options
```bash
ecs-watcher --help
Usage of ecs-watcher:
  -check.interval duration
        The interval for checking the cluster (default 5s)
  -cluster string
        The target cluster name
  -debug
        Run in debug mode
  -gc.interval duration
        The minimum interval for garbage collection of unhealthy targets (default 2s)
  -gc.step.percent int
        The step percent of total unhealthy targets when cleaning (default 20)
  -region string
        The AWS region of the cluster
  -unhealthy.after duration
        The duration that a target needs to be unhealthy to declare as unhealthy (default 1m0s)
  -unhealthy.tag string
        The tag used to mark unhealty labels key:value form (default "unhealthy:true")

```

## Install

### from Source
TODO

### Docker
TODO

### Release
TODO
