package main

import (
	"errors"
	"testing"
	"time"
)

type testChecker struct {
	checkCounter     int
	markCounter      int
	checkReturnError bool
	markReturnError  bool
}

func (t *testChecker) Check() error {
	t.checkCounter++
	if t.checkReturnError {
		return errors.New("")
	}
	return nil
}

func (t *testChecker) Mark() error {
	t.markCounter++
	if t.markReturnError {
		return errors.New("")
	}
	return nil
}

func TestWatcherNoChecker(t *testing.T) {
	w := &Watcher{}
	err := w.Run()
	if err == nil {
		t.Errorf("Watcher run without checker should give error")
	}
}

func TestWatcherRunsInInterval(t *testing.T) {
	wait := 1*time.Second + 10*time.Millisecond
	interval := 50 * time.Millisecond
	c := &testChecker{}

	w := &Watcher{
		interval: interval,
		checker:  c,
	}

	// Run in background
	go w.Run()

	// Wait a second
	time.Sleep(wait)

	// check if check and mark ran the expected times
	expected := wait / interval

	if c.checkCounter != int(expected) {
		t.Errorf("Checker check ran times is wrong. Expected: %d, got: %d", expected, c.checkCounter)
	}

	if c.markCounter != int(expected) {
		t.Errorf("Checker mark ran times is wrong. Expected: %d, got: %d", expected, c.markCounter)
	}

}

func TestWatcherRunsInIntervalWithCheckError(t *testing.T) {
	wait := 1*time.Second + 10*time.Millisecond
	interval := 200 * time.Millisecond
	c := &testChecker{checkReturnError: true}

	w := &Watcher{
		interval: interval,
		checker:  c,
	}

	// Run in background
	go w.Run()

	// Wait a second
	time.Sleep(wait)

	// check if check and mark ran the expected times
	expected := wait / interval

	if c.checkCounter != int(expected) {
		t.Errorf("Checker check ran times is wrong. Expected: %d, got: %d", expected, c.checkCounter)
	}

	if c.markCounter != 0 {
		t.Errorf("Checker mark ran times is wrong. Expected: %d, got: %d", 0, c.markCounter)
	}
}

func TestWatcherRunsInIntervalWithMarkError(t *testing.T) {
	wait := 1*time.Second + 10*time.Millisecond
	interval := 200 * time.Millisecond
	c := &testChecker{markReturnError: true}

	w := &Watcher{
		interval: interval,
		checker:  c,
	}

	// Run in background
	go w.Run()

	// Wait a second
	time.Sleep(wait)

	// check if check and mark ran the expected times
	expected := wait / interval

	if c.checkCounter != int(expected) {
		t.Errorf("Checker check ran times is wrong. Expected: %d, got: %d", expected, c.checkCounter)
	}

	if c.markCounter != int(expected) {
		t.Errorf("Checker mark ran times is wrong. Expected: %d, got: %d", int(expected), c.markCounter)
	}
}
