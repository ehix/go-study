package dberrors

type ConflictError struct{}

// Create a new instance of ConflictError so DB logic doesn't leak into web layer.
func (e *ConflictError) Error() string {
	return "attempted to create a.record with an existing key"
}
