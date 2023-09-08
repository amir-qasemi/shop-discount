package order

import (
	"errors"

	"github.com/amir-qasemi/shop-discount/internal/user"
)

// Service for retrieving orders.
// This could be a gRPC consumer of another microservice or
// a service implemented in this same module
type Service interface {
	// GetUserOrders get all orders of the given user
	GetUserOrders(user *user.User) ([]Order, error)
	// Get all orders of the given user which are in the given orderStatus
	GetUserOrdersInStatus(user *user.User, orderStatus []OrderStatus) ([]Order, error)
}

// DummyService A sample service for mocking purposes
type DummyService struct {
}

func (o *DummyService) GetUserOrdersInStatus(user *user.User, orderStatus []OrderStatus) ([]Order, error) {
	if user.Username == "NewUser" {
		return []Order{}, nil
	} else {
		return nil, errors.New("Problem in fetching order: Mocking method for given conditions not implemented")
	}
}

func (o *DummyService) GetUserOrders(user *user.User) ([]Order, error) {
	if user.Username == "NewUser" {
		return []Order{}, nil
	} else {
		return nil, errors.New("Problem in fetching order: Mocking method for given conditions not implemented")
	}
}
