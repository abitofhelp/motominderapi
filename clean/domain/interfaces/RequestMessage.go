// Package interfaces contains contracts for entities and other objects.
package interfaces

// RequestMessage defines the base contract for all request messages.
type RequestMessage interface {
	// Validate verifies that a RequestMessage's fields contain valid data.
	// Returns nil if the RequestMessage contains valid data, otherwise an error.
	Validate() error
}
