package cart

// Service for retrieving the cart.
// This could be a gRPC consumer of another microservice or
// a service implemented in this same module
type Service interface {
	// Get the cart by id or an error if the cart does not exist
	GetCartById(string) (*Cart, error)
}
