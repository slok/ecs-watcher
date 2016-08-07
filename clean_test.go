package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/golang/mock/gomock"

	awsMock "github.com/slok/ecs-watcher/mock/aws"
	"github.com/slok/ecs-watcher/mock/aws/sdk"
)

func TestAgentCheckerListContainerInstancesError(t *testing.T) {

	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockECSCli := sdk.NewMockECSAPI(ctrl)

	// Set our mock desired result
	awsMock.MockListContainerInstancesPagesError(t, mockECSCli)

	a := &AgentChecker{
		clusterName:      "test",
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
	}
	a.ecsCli = mockECSCli

	if err := a.Check(); err == nil {
		t.Errorf("Check should give an error, it didn't")
	}
}

func TestAgentCheckerListContainerInstancesZero(t *testing.T) {
	quantity := 0

	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockECSCli := sdk.NewMockECSAPI(ctrl)

	// Set our mock desired result
	awsMock.MockListContainerInstancesPagesQ(t, mockECSCli, quantity)

	a := &AgentChecker{
		clusterName:      "test",
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
	}
	a.ecsCli = mockECSCli

	err := a.Check()
	if err != nil {
		t.Errorf("Check should'n give an error: %s", err)
	}

	if len(a.unhealthies) != 0 {
		t.Errorf("Unhealthy instances should be 0, got %d", len(a.unhealthies))
	}
}

func TestAgentCheckerDescribeContainerInstancesError(t *testing.T) {
	quantity := 10

	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockECSCli := sdk.NewMockECSAPI(ctrl)

	// Set our mock desired result
	awsMock.MockListContainerInstancesPagesQ(t, mockECSCli, quantity)
	awsMock.MockDescribeContainerInstancesError(t, mockECSCli)

	a := &AgentChecker{
		clusterName:      "test",
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
	}
	a.ecsCli = mockECSCli

	err := a.Check()
	if err == nil {
		t.Errorf("Check should give an error, it didn't")
	}
}

func TestAgentCheckerDescribeContainerInstancesHealthiness(t *testing.T) {
	tests := []struct {
		healthy   int
		unhealthy int
	}{
		{10, 20},
		{10, 0},
		{0, 0},
		{0, 30},
		{200, 500},
	}

	for _, test := range tests {
		quantity := test.healthy + test.unhealthy

		// Create mock for AWS API
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockECSCli := sdk.NewMockECSAPI(ctrl)

		// Set our mock desired result
		awsMock.MockListContainerInstancesPagesQ(t, mockECSCli, quantity)
		awsMock.MockDescribeContainerInstancesHealthyUnhealthyQ(t, mockECSCli, test.healthy, test.unhealthy)

		a := &AgentChecker{
			clusterName:      "test",
			unhealthies:      make(map[string]*unhealthyInstance),
			unhealthiesMutex: &sync.Mutex{},
		}
		a.ecsCli = mockECSCli

		err := a.Check()
		if err != nil {
			t.Errorf("-%+v\n- Check shouldn't give an error: %s", test, err)
		}

		if len(a.unhealthies) != test.unhealthy {
			t.Errorf("-%+v\n- Wrong number of unhealthy instances; got: %d, want: %d", test, len(a.unhealthies), test.unhealthy)
		}
	}
}

func TestAgentCheckerDescribeContainerInstancesUnhealthyKeepOld(t *testing.T) {
	tests := []struct {
		unhealthyCalls [][]string
		wantTotal      []int
	}{
		{
			unhealthyCalls: [][]string{
				[]string{"a", "b", "c"},
				[]string{"a", "d", "e"}, // b & c have sanitized
				[]string{"f"},           // a, d and e have sanitized
			},
		},
		{
			unhealthyCalls: [][]string{
				[]string{"a", "b", "c", "d", "e", "f", "g"},
				[]string{"a", "b", "c", "d", "e", "f", "g"},
				[]string{"a", "b", "c", "d", "e", "f", "h"},
			},
		},
		{
			unhealthyCalls: [][]string{
				[]string{"a"},
				[]string{"b"},
				[]string{"c"},
			},
		},
		{
			unhealthyCalls: [][]string{
				[]string{"a"},
				[]string{"b"},
				[]string{},
			},
		},
	}

	for _, test := range tests {

		// Create mock for AWS API
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockECSCli := sdk.NewMockECSAPI(ctrl)

		a := &AgentChecker{
			clusterName:      "test",
			unhealthies:      make(map[string]*unhealthyInstance),
			unhealthiesMutex: &sync.Mutex{},
		}
		a.ecsCli = mockECSCli

		previousResult := map[string]time.Time{}

		// Start multiple calls seing the result
		for _, c := range test.unhealthyCalls {
			// Set our mock desired result
			awsMock.MockListContainerInstancesPagesQ(t, mockECSCli, len(c))
			awsMock.MockDescribeContainerInstancesUnhealthies(t, mockECSCli, c...)

			err := a.Check()
			if err != nil {
				t.Errorf("-%+v\n- Check shouldn't give an error: %s", test, err)
			}
			// Check length
			if len(a.unhealthies) != len(c) {
				t.Errorf("-%+v\n- Wrong number of unhealthy instances; got: %d, want: %d", test, len(a.unhealthies), len(c))
			}

			// Check the old timestamp is not new
			for k, v := range a.unhealthies {
				// Check if we persist correctly the timestamp from previous calls
				oldTs, ok := previousResult[k]
				if ok {
					if v.started != oldTs {
						t.Errorf("-%+v\n-%v\n- Wrong timestamp, the timestamp of the same unhealthy instance after multiple calls should be the same; got: %s, want: %s", test, k, oldTs, v.started)
					}
				}
				// store this call to check afterwards on the next call
				previousResult[k] = v.started
			}

		}
	}
}

func TestAgentCheckerMarkZeroUnhealthies(t *testing.T) {

	a := &AgentChecker{
		clusterName:      "test",
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
	}

	err := a.Mark()
	if err != nil {
		t.Errorf("Check should'n give an error: %s", err)
	}

	if len(a.unhealthies) != 0 {
		t.Errorf("Unhealthy instances should be 0, got %d", len(a.unhealthies))
	}
}

func TestAgentCheckerMarkError(t *testing.T) {
	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEC2Cli := sdk.NewMockEC2API(ctrl)

	// Set our mock desired result
	awsMock.MockCreateTagsError(t, mockEC2Cli)

	a := &AgentChecker{
		clusterName:      "test",
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
	}
	a.ec2Cli = mockEC2Cli

	// Add one in order to continue with the flow of calling AWS API
	a.unhealthies["a"] = &unhealthyInstance{
		instance: &ecs.ContainerInstance{},
		started:  time.Now().UTC(),
	}

	err := a.Mark()
	if err == nil {
		t.Errorf("Mark should give an error, it didn't")
	}
}

func TestAgentCheckerMarkAfterTime(t *testing.T) {
	tests := []struct {
		markAfter time.Duration
		started   time.Time

		shouldMark bool
	}{
		{30 * time.Second, time.Now().UTC().Add(-29 * time.Second), false},
		{30 * time.Second, time.Now().UTC().Add(-31 * time.Second), true},
		{10 * time.Minute, time.Now().UTC().Add(-5 * time.Minute), false},
		{1 * time.Second, time.Now().UTC().Add(-1 * time.Second), true},
	}

	for _, test := range tests {
		// Create mock for AWS API
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockEC2Cli := sdk.NewMockEC2API(ctrl)

		// Set our mock desired result
		marked := map[string]string{}
		awsMock.MockCreateTags(t, mockEC2Cli, marked)

		a := &AgentChecker{
			clusterName:      "test",
			unhealthies:      make(map[string]*unhealthyInstance),
			unhealthiesMutex: &sync.Mutex{},
			markAfter:        test.markAfter,
		}
		a.ec2Cli = mockEC2Cli

		a.unhealthies["test"] = &unhealthyInstance{
			instance: &ecs.ContainerInstance{},
			started:  test.started,
		}

		err := a.Mark()
		if err != nil {
			t.Errorf("-%+v\n  -Mark shouldn't give an error: %s", test, err)
		}

		// Check if tagged
		_, ok := a.unhealthies["test"]
		if test.shouldMark {
			if ok {
				t.Errorf("-%+v\n  -After marking, instance shouldn't be in unhealthy ones", test)
			}

			if _, ok = marked["test"]; !ok {
				t.Errorf("-%+v\n  -After marking, instance should be marked by the API, it isn't", test)
			}
		}

		if !test.shouldMark && !ok {
			t.Errorf("-%+v\n  -After not marking, instance should continue in unhealthy ones", test)
		}

	}
}

func TestAgentCheckerMarkMultiple(t *testing.T) {
	quantity := 100

	// Create mock for AWS API
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEC2Cli := sdk.NewMockEC2API(ctrl)

	// Set our mock desired result
	marked := map[string]string{}
	awsMock.MockCreateTags(t, mockEC2Cli, marked)

	a := &AgentChecker{
		clusterName:      "test",
		unhealthies:      make(map[string]*unhealthyInstance),
		unhealthiesMutex: &sync.Mutex{},
		markAfter:        30 * time.Second,
		markTag:          MarkTag{key: "key", value: "value"},
	}
	a.ec2Cli = mockEC2Cli

	for i := 0; i < quantity; i++ {
		a.unhealthies[fmt.Sprintf("i-%d", i)] = &unhealthyInstance{
			instance: &ecs.ContainerInstance{},
			started:  time.Now().UTC().Add(-1 * time.Minute),
		}
	}

	err := a.Mark()
	if err != nil {
		t.Errorf("Mark shouldn't give an error: %s", err)
	}

	// Check if tagged
	if len(a.unhealthies) != 0 {
		t.Errorf("After marking them they should be deleted, this should be empty, got: %d", len(a.unhealthies))
	}

	if len(marked) != quantity {
		t.Errorf("Wrong number of instances marked, got: %d, want: %d", len(marked), quantity)
	}

	for _, tag := range marked {
		if tag != "key:value" {
			t.Errorf("Wrong tag on instance: %s", tag)
		}
	}
}
