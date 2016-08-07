package main

import (
	"sync"
	"testing"
	"time"

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
