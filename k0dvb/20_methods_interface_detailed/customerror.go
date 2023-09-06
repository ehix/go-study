package main

import "fmt"

// Custom error type, has error in it.
type errFoo struct {
	err  error
	path string
}

// This makes errFoo compatible with type `error` (satisfies it's interface).
func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

// Returns concrete ptr to concrete type errFoo, problem.
func XYZ(a int) *errFoo {
	// Not nil in the interface sence, but ptr to errFoo sence.
	return nil
}

// Returns returns a nil interface, excepted.
func XYZalt(a int) error {
	return nil
}

func main() {
	// err := XYZ(1) // err would be of type *errFoo, not an error interface.
	var brokenErr error = XYZ(1) // BAD: interface gets a nil concrete ptr
	if brokenErr != nil {
		// The interface isn't nil, even though it has a nil ptr inside.
		fmt.Println("Oops")
	} else {
		fmt.Println("OK!")
	}

	var fixedErr error = XYZalt(1)
	if fixedErr != nil {
		fmt.Println("Oops")
	} else {
		fmt.Println("OK!")
	}
}
