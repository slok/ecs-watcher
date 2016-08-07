package main

// Checker is an interface that represents a cluster checker
type Checker interface {
	// Will check if the cluster is ok
	Check() error

	// Mark will mark the spoted unhealthy checked stuff
	Mark() error
}

// Cleaner interface represents the one that will take the action of cleaning marked targets
type Cleaner interface {
	Clean() error
}

//MarkTag represents the marking tag
type MarkTag struct {
	key   string
	value string
}
