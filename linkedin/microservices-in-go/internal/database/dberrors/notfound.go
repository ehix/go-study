package dberrors

import "fmt"

// Create a new instance of NotFoundError so DB logic doesn't leak into web layer.
type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("unable to find %s with id %s", e.Entity, e.ID)
}
