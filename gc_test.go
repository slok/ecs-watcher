package main

import (
	"errors"
	"testing"
	"time"
)

type testCleaner struct {
	cleanCounter     int
	cleanReturnError bool
}

func (t *testCleaner) Clean() error {
	t.cleanCounter++
	if t.cleanReturnError {
		return errors.New("")
	}
	return nil
}

func TestGCNoCleaner(t *testing.T) {
	gc := &GC{}
	err := gc.Run()
	if err == nil {
		t.Errorf("Garbage collector run without cleaner should give error")
	}
}

func TestGCRunsInInterval(t *testing.T) {
	wait := 1*time.Second + 10*time.Millisecond
	interval := 50 * time.Millisecond
	c := &testCleaner{}

	gc := &GC{
		interval: interval,
		cleaner:  c,
	}

	// Run in background
	go gc.Run()

	// Wait a second
	time.Sleep(wait)

	// check if clean ran the expected times
	expected := wait / interval

	if c.cleanCounter != int(expected) {
		t.Errorf("Cleaner clean ran times is wrong. Expected: %d, got: %d", expected, c.cleanCounter)
	}
}

func TestGCRunsInIntervalWithCleanError(t *testing.T) {
	wait := 1*time.Second + 10*time.Millisecond
	interval := 200 * time.Millisecond
	c := &testCleaner{cleanReturnError: true}

	gc := &GC{
		interval: interval,
		cleaner:  c,
	}

	// Run in background
	go gc.Run()

	// Wait a second
	time.Sleep(wait)

	// check if clean ran the expected times
	expected := wait / interval

	if c.cleanCounter != int(expected) {
		t.Errorf("Cleaner clean ran times is wrong. Expected: %d, got: %d", expected, c.cleanCounter)
	}

}
