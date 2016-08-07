package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		args    []string
		correct bool
	}{
		{[]string{"--debug"}, false},
		{[]string{"-debug", "--region", "eu-west-1"}, false},
		{[]string{"--region", "eu-west-1", "-cluster", "test"}, true},
		{[]string{"--region", "eu-west-1", "-cluster", "test", "-unhealthy.tag", "key-value"}, false},
		{[]string{"--region", "eu-west-1", "-cluster", "test", "-unhealthy.tag", "key:value"}, true},
		{[]string{"--region", "eu-west-1", "-cluster", "test", "check.interval", "1t"}, false},
	}

	for _, test := range tests {
		err := parse(test.args)
		if err != nil && test.correct {
			t.Errorf("- %+v\n Shouldn't give an error: %s", test, err)
		}

		if err == nil && !test.correct {
			t.Errorf("- %+v\n Shouldt give an error, it didn't", test)
		}

	}
}
