package aws

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/golang/mock/gomock"

	"github.com/slok/ecs-watcher/mock/aws/sdk"
)

// MockListContainerInstancesPagesError will error when calling
func MockListContainerInstancesPagesError(t *testing.T, mockMatcher *sdk.MockECSAPI) {
	logrus.Warningf("Mocking AWS iface: ListContainerInstancesPages")

	err := errors.New("Wrong!")

	mockMatcher.EXPECT().ListContainerInstancesPages(gomock.Any(), gomock.Any()).AnyTimes().Return(err)
}

// MockListContainerInstancesPagesQ will return q arns when calling
func MockListContainerInstancesPagesQ(t *testing.T, mockMatcher *sdk.MockECSAPI, q int) {
	logrus.Warningf("Mocking AWS iface: ListContainerInstancesPages")

	var err error

	mockMatcher.EXPECT().ListContainerInstancesPages(gomock.Any(), gomock.Any()).Do(
		func(input *ecs.ListContainerInstancesInput, fn func(p *ecs.ListContainerInstancesOutput, lastPage bool) (shouldContinue bool)) {
			cs := make([]*string, q)
			for i := 0; i < q; i++ {
				cs[i] = aws.String(fmt.Sprintf("arn-%d", i))
			}
			resp := &ecs.ListContainerInstancesOutput{
				ContainerInstanceArns: cs,
			}
			fn(resp, true)
		}).AnyTimes().Return(err)
}

// MockDescribeContainerInstancesError will error when calling
func MockDescribeContainerInstancesError(t *testing.T, mockMatcher *sdk.MockECSAPI) {
	logrus.Warningf("Mocking AWS iface: DescribeContainerInstances")

	err := errors.New("Wrong!")

	mockMatcher.EXPECT().DescribeContainerInstances(gomock.Any()).AnyTimes().Return(nil, err)
}

// MockDescribeContainerInstancesHealthyUnhealthyQ will return the desired healthy & unhealthy instances
func MockDescribeContainerInstancesHealthyUnhealthyQ(t *testing.T, mockMatcher *sdk.MockECSAPI, healthy, unhealthy int) {
	logrus.Warningf("Mocking AWS iface: DescribeContainerInstances")

	var err error

	cs := make([]*ecs.ContainerInstance, healthy+unhealthy)

	// Create healthy ones
	for i := 0; i < healthy; i++ {
		cs[i] = &ecs.ContainerInstance{
			Ec2InstanceId:  aws.String(fmt.Sprintf("i-%d", i)),
			AgentConnected: aws.Bool(true),
		}
	}

	// Create unhealthy ones
	for i := healthy; i < healthy+unhealthy; i++ {
		cs[i] = &ecs.ContainerInstance{
			Ec2InstanceId:  aws.String(fmt.Sprintf("i-%d", i)),
			AgentConnected: aws.Bool(false),
		}
	}

	resp := &ecs.DescribeContainerInstancesOutput{
		ContainerInstances: cs,
	}
	mockMatcher.EXPECT().DescribeContainerInstances(gomock.Any()).AnyTimes().Return(resp, err)
}

// MockDescribeContainerInstancesUnhealthies will return the ddesired unhealthy instances with the ids
func MockDescribeContainerInstancesUnhealthies(t *testing.T, mockMatcher *sdk.MockECSAPI, unhealthyIds ...string) {
	logrus.Warningf("Mocking AWS iface: DescribeContainerInstances")

	var err error

	cs := make([]*ecs.ContainerInstance, len(unhealthyIds))
	for i, u := range unhealthyIds {
		cs[i] = &ecs.ContainerInstance{
			Ec2InstanceId:  aws.String(u),
			AgentConnected: aws.Bool(false),
		}
	}

	resp := &ecs.DescribeContainerInstancesOutput{
		ContainerInstances: cs,
	}
	// Only one time
	mockMatcher.EXPECT().DescribeContainerInstances(gomock.Any()).Return(resp, err)
}
