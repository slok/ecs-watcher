package main

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

const instanceStateRunningCode = "16"

// Killer will clean unhealthy imaes that are tagged
type Killer struct {
	ec2Cli     ec2iface.EC2API
	ec2WaitCli *ec2.EC2
	session    *session.Session

	// the step percent of cleaning instances
	step int

	// The tag that marked instnaces to clean have
	markTag MarkTag

	// Should we wait to instance terminated status?
	waitTerminate bool
}

// NewKiller creates a new killer
func NewKiller(awsRegion string, stepPercent int, mtag string) (*Killer, error) {
	k := &Killer{
		step:          stepPercent,
		waitTerminate: true,
	}

	// Set the tag
	splTag := strings.Split(mtag, ":")
	k.markTag = MarkTag{splTag[0], splTag[1]}

	// Create AWS session
	s := session.New(&aws.Config{Region: aws.String(awsRegion)})
	if s == nil {
		return nil, fmt.Errorf("error creating aws session")
	}
	k.session = s

	// Create the AWS EC2 client
	k.ec2Cli = ec2.New(s)

	// Create the wait client, the API interface doesn't implement the waiters
	k.ec2WaitCli = ec2.New(s)

	return k, nil
}

// Clean will hunt and kill unhealthy instances
func (k *Killer) Clean() error {
	// Get all the marked instances
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String(fmt.Sprintf("tag:%s", k.markTag.key)),
				Values: []*string{aws.String(k.markTag.value)},
			},
			{
				Name:   aws.String("instance-state-code"),
				Values: []*string{aws.String(instanceStateRunningCode)},
			},
		},
	}

	var instances []*ec2.Instance

	err := k.ec2Cli.DescribeInstancesPages(params,
		func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
			for _, r := range page.Reservations {
				for _, i := range r.Instances {
					instances = append(instances, i)
				}
			}
			return true
		})

	if err != nil {
		return err
	}
	if len(instances) == 0 {
		logrus.Debugf("No targets to kill")
		return nil
	}
	logrus.Debugf("Killing targets: %d", len(instances))

	// Get the number of instances per step
	n := k.step * len(instances) / 100
	if n == 0 {
		n = 1
	}
	logrus.Infof("Start killing in batches of %d", n)

	// Start killing them in steps and wait until it was terminated
	for i := 0; i < len(instances); i = i + n {
		var targets []*ec2.Instance
		if i+n > len(instances) {
			targets = instances[i:]
		} else {
			targets = instances[i : i+n]
		}

		// Kill
		ids := make([]*string, len(targets))
		for it, t := range targets {
			ids[it] = t.InstanceId
		}

		if len(targets) == 0 {
			logrus.Debugf("Nothing to kill")
			return nil
		}

		params := &ec2.TerminateInstancesInput{
			InstanceIds: ids,
		}

		// Finish him!
		_, err = k.ec2Cli.TerminateInstances(params)
		if err != nil {
			return err
		}

		// Wait if wanted
		if k.waitTerminate {
			paramsWait := &ec2.DescribeInstancesInput{
				InstanceIds: ids,
			}
			err = k.ec2WaitCli.WaitUntilInstanceTerminated(paramsWait)
			if err != nil {
				return err
			}
		}
		logrus.Infof("Killed %d targets", len(ids))
	}

	return nil
}
