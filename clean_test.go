package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"

	awsMock "github.com/slok/ecs-watcher/mock/aws"
	"github.com/slok/ecs-watcher/mock/aws/sdk"
)

func TestKillerDesribeInstancesError(t *testing.T) {
	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEC2Cli := sdk.NewMockEC2API(ctrl)

	// Set our mock desired result
	awsMock.MockDescribeInstancesPagesError(t, mockEC2Cli)

	k := &Killer{
		markTag: MarkTag{"key", "value"},
	}
	k.ec2Cli = mockEC2Cli

	err := k.Clean()

	if err == nil {
		t.Errorf("Clean should give an error, it didn't")
	}
}

func TestKillerTerminateZeroInstances(t *testing.T) {

	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEC2Cli := sdk.NewMockEC2API(ctrl)

	// Set our mock desired result
	awsMock.MockDescribeInstancesPagesQ(t, mockEC2Cli, 0, 0)
	awsMock.MockTerminateInstancesError(t, mockEC2Cli)

	k := &Killer{
		markTag: MarkTag{"key", "value"},
	}
	k.ec2Cli = mockEC2Cli

	err := k.Clean()

	if err != nil {
		t.Errorf("Clean shouldn't give an error: %s", err)
	}

}

func TestKilleTerminateMultiple(t *testing.T) {
	tests := []struct {
		unhealthyQ int
		step       int

		wantCalls int
	}{
		{20, 10, 10},
		{0, 10, 0},
		{20, 50, 2},
		{20, 100, 1},
		{500, 1, 100},
		{1000, 15, 7},
		{47, 0, 47},
		{21, 10, 11},
	}

	for _, test := range tests {
		// Create mock for AWS API
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockEC2Cli := sdk.NewMockEC2API(ctrl)

		// Set our mock desired result
		terminatedCalls := []map[string]*ec2.InstanceState{}
		awsMock.MockDescribeInstancesPagesQ(t, mockEC2Cli, test.unhealthyQ, 0)
		awsMock.MockTerminateInstances(t, mockEC2Cli, &terminatedCalls)

		k := &Killer{
			markTag:       MarkTag{"key", "value"},
			step:          test.step,
			waitTerminate: false,
		}
		k.ec2Cli = mockEC2Cli

		err := k.Clean()

		if err != nil {
			t.Errorf("%+v\n- Clean shouldn't give an error: %s", test, err)
		}

		if len(terminatedCalls) != test.wantCalls {
			t.Errorf("%+v\n- Wrong number of calls to terminate instances: got: %d, want: %d", test, len(terminatedCalls), test.wantCalls)
		}

		totalSum := 0
		for _, c := range terminatedCalls {
			totalSum += len(c)
		}

		if totalSum != test.unhealthyQ {
			t.Errorf("%+v\n- Wrong number of terminated instances; got: %d, want: %d", test, totalSum, test.unhealthyQ)
		}

	}
}
