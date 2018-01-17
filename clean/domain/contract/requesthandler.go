// Package contract contains contracts for entities and other objects.
package contract

// RequestHandler is the default handler for a requests.
type RequestHandler interface {
	Handle(requestMessage RequestMessage) (ResponseMessage, error)
}
