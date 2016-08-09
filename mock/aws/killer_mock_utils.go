package aws

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/slok/ecs-watcher/mock/aws/sdk"
)

// MockDescribeInstancesPagesError will return an error
func MockDescribeInstancesPagesError(t *testing.T, mockMatcher *sdk.MockEC2API) {
	logrus.Warningf("Mocking AWS iface: DescribeInstancesPages")

	err := errors.New("")

	mockMatcher.EXPECT().DescribeInstancesPages(gomock.Any(), gomock.Any()).AnyTimes().Return(err)
}

// MockDescribeInstancesPagesQ will return q instances when calling
func MockDescribeInstancesPagesQ(t *testing.T, mockMatcher *sdk.MockEC2API, runningQ, notRunningQ int) {
	logrus.Warningf("Mocking AWS iface: DescribeInstancesPages")

	var err error

	mockMatcher.EXPECT().DescribeInstancesPages(gomock.Any(), gomock.Any()).Do(
		func(input *ec2.DescribeInstancesInput, fn func(p *ec2.DescribeInstancesOutput, lastPage bool) (shouldContinue bool)) {
			instances := make([]*ec2.Instance, runningQ+notRunningQ)
			// Stopped ones
			for i := 0; i < notRunningQ; i++ {
				instances[i] = &ec2.Instance{
					InstanceId: aws.String(fmt.Sprintf("i-%d", i)),
					State:      &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameStopped)},
				}
			}
			// Running ones
			for i := notRunningQ; i < notRunningQ+runningQ; i++ {
				instances[i] = &ec2.Instance{
					InstanceId: aws.String(fmt.Sprintf("i-%d", i)),
					State:      &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameRunning)},
				}
			}

			resp := &ec2.DescribeInstancesOutput{
				Reservations: []*ec2.Reservation{
					&ec2.Reservation{Instances: instances},
				},
			}
			fn(resp, true)
		}).AnyTimes().Return(err)
}

// MockTerminateInstancesError will return error
func MockTerminateInstancesError(t *testing.T, mockMatcher *sdk.MockEC2API) {
	logrus.Warningf("Mocking AWS iface: TerminateInstances")

	err := errors.New("")

	mockMatcher.EXPECT().TerminateInstances(gomock.Any()).AnyTimes().Return(nil, err)
}

// MockTerminateInstances will set on the instances the terminated received ones
func MockTerminateInstances(t *testing.T, mockMatcher *sdk.MockEC2API, instances *[]map[string]*ec2.InstanceState) {
	var err error
	mockMatcher.EXPECT().TerminateInstances(gomock.Any()).Do(
		func(input *ec2.TerminateInstancesInput) {
			m := make(map[string]*ec2.InstanceState)
			*instances = append(*instances, m)
			for _, i := range input.InstanceIds {
				m[aws.StringValue(i)] = &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameRunning)}
			}
		}).AnyTimes().Return(nil, err)

}
