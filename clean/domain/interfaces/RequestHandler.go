// Package interfaces contains contracts for entities and other objects.
package interfaces

// RequestHandler is the default handler for a requests.
type RequestHandler interface {
	Handle(requestMessage RequestMessage) (ResponseMessage, error)
}
