package user

// Service for retrieving the users.
// This could be a gRPC consumer of another microservice or
// a service implemented in this same module
type Service interface {
	GetUser(string) (*User, error)
}

// DummyService A sample service for mocking purposes
type DummyService struct {
}

func (s *DummyService) GetUser(username string) (*User, error) {
	return &User{Username: username}, nil
}
