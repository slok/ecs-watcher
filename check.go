package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

const (
	checkMaxAWSAPIResult = 50
)

type unhealthyInstance struct {
	instance *ecs.ContainerInstance
	started  time.Time
}

// AgentChecker will check the agent status on the ECS cluster instances
type AgentChecker struct {
	ecsCli  ecsiface.ECSAPI
	ec2Cli  ec2iface.EC2API
	session *session.Session

	// the name of the cluster
	clusterName string

	// unhealthy instances
	unhealthies      map[string]*unhealthyInstance
	unhealthiesMutex *sync.Mutex

	// The tag to mark tge unhealthy instances
	markTag MarkTag
}

// NewAgentChecker creates an AgentChecker
func NewAgentChecker(clusterName string, awsRegion string, tag string) (*AgentChecker, error) {
	a := &AgentChecker{
		clusterName:      clusterName,
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
	}

	// Set the tag
	splTag := strings.Split(tag, ":")
	a.markTag = MarkTag{splTag[0], splTag[1]}

	// Create AWS session
	s := session.New(&aws.Config{Region: aws.String(awsRegion)})
	if s == nil {
		return nil, fmt.Errorf("error creating aws session")
	}
	a.session = s

	// Create AWS ecs client
	ec := ecs.New(s)
	a.ecsCli = ec

	// Create the AWS EC2 client
	a.ec2Cli = ec2.New(s)

	return a, nil
}

// Check will check if the agent is connected in each instance
func (a *AgentChecker) Check() error {

	logrus.Debugf("Getting cluster instance arns")

	// Get the container instance ARNs
	lparams := &ecs.ListContainerInstancesInput{
		Cluster:    aws.String(a.clusterName),
		MaxResults: aws.Int64(checkMaxAWSAPIResult),
	}
	var arns []*string
	err := a.ecsCli.ListContainerInstancesPages(lparams,
		func(page *ecs.ListContainerInstancesOutput, lastPage bool) bool {
			// Append the arns
			arns = append(arns, page.ContainerInstanceArns...)
			return true
		})

	if err != nil {
		return err
	}
	if len(arns) == 0 {
		logrus.Warningf("No container instances present")
		return nil
	}

	logrus.Debugf("Got %d container instance ARNs", len(arns))

	// Check the status of the container instances
	dparams := &ecs.DescribeContainerInstancesInput{
		ContainerInstances: arns,
		Cluster:            aws.String(a.clusterName),
	}
	resp, err := a.ecsCli.DescribeContainerInstances(dparams)
	if err != nil {
		return err
	}

	// Use this as counter, maybe the older unhealty ones are in the process of
	// removal, so we can't use the unhelty total as the cluster unhealthy total number
	var unhCount int
	a.unhealthiesMutex.Lock()
	// Save the unhealthy ones
	for _, ci := range resp.ContainerInstances {
		// if ok don't do nothing
		if aws.BoolValue(ci.AgentConnected) {
			continue
		}
		logrus.Warning("%+v", ci)

		// TODO: Check status of the instance?

		// The instance seems unhealthy because of the agent, check if is already there
		if _, ok := a.unhealthies[aws.StringValue(ci.Ec2InstanceId)]; !ok {
			unhCount++
			ui := &unhealthyInstance{
				instance: ci,
				started:  time.Now().UTC(),
			}
			a.unhealthies[aws.StringValue(ci.Ec2InstanceId)] = ui
		}
	}
	a.unhealthiesMutex.Unlock()

	if unhCount > 0 {
		logrus.Infof("new %d of %d are unhealthy", unhCount, len(arns))
	} else {
		logrus.Debugf("No new unhealthy instances")
	}
	logrus.Infof("%d total unhealthy", len(a.unhealthies))

	return nil
}

// Mark will mark them as unhealthy
func (a *AgentChecker) Mark() error {
	a.unhealthiesMutex.Lock()
	defer a.unhealthiesMutex.Unlock()

	if len(a.unhealthies) == 0 {
		logrus.Debugf("Skipping marking, no unhealthy instances")
		return nil
	}

	var resources []*string

	for id := range a.unhealthies {
		resources = append(resources, aws.String(id))
	}

	// mark all the unhealthy images
	params := &ec2.CreateTagsInput{
		Resources: resources,
		Tags: []*ec2.Tag{
			{Key: aws.String(a.markTag.key), Value: aws.String(a.markTag.value)},
		},
	}
	_, err := a.ec2Cli.CreateTags(params)
	if err != nil {
		return err
	}

	logrus.Infof("Marked %d", len(a.unhealthies))
	// Reset the unhealthies, they are already marked (we are save because of the mutex)
	a.unhealthies = make(map[string]*unhealthyInstance)

	return nil
}
