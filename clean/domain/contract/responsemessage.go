// Package contract contains contracts for entities and other objects.
package contract

// ResponseMessage defines the base contract for all response messages.
type ResponseMessage interface {
	// Validate verifies that a ResponseMessage's fields contain valid data.
	// Returns nil if the ResponseMessage contains valid data, otherwise an error.
	Validate() error
}
