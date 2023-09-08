package cart

import "github.com/amir-qasemi/shop-discount/internal/product"

// Service for retrieving the cart.
// This could be a gRPC consumer of another microservice or
// a service implemented in this same module
type Service interface {
	// Get the cart by id or an error if the cart does not exist
	GetCartById(string) (*Cart, error)
}

// DummyService A sample service for mocking purposes
type DummyService struct {
}

func (s *DummyService) GetCartById(id string) (*Cart, error) {
	line1 := &Line{
		Product: product.SpecializedProduct{
			Product:    product.Product{Id: "1"},
			Price:      100,
			Attributes: make(map[string]string),
		},
		Num:                1,
		PriceAfterDiscount: 100,
	}

	line2 := &Line{
		Product: product.SpecializedProduct{
			Product:    product.Product{Id: "2"},
			Price:      200,
			Attributes: make(map[string]string),
		},
		Num:                1,
		PriceAfterDiscount: 200,
	}

	line3 := &Line{
		Product: product.SpecializedProduct{
			Product:    product.Product{Id: "3"},
			Price:      300,
			Attributes: make(map[string]string),
		},
		Num:                2,
		PriceAfterDiscount: 600,
	}

	lines := []*Line{line1, line2, line3}

	return &Cart{
		Id:    id,
		Lines: lines,
	}, nil
}
