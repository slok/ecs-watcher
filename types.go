package main

// Checker is an interface that represents a cluster checker
type Checker interface {
	// Will check if the cluster is ok
	Check() error

	// Mark will mark the spoted unhealthy checked stuff
	Mark() error
}
