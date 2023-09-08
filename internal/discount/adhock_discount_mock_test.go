package discount

import (
	"sync"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/product"
	"github.com/amir-qasemi/shop-discount/internal/user"
	"github.com/stretchr/testify/mock"
)

type MockedRepository struct {
	mock.Mock
}

func (r *MockedRepository) GetDiscountByCode(code string) (Discount, error) {
	args := r.Called(code)
	return args.Get(0).(Discount), args.Error(1)
}

func (r *MockedRepository) GetDiscountByUsageId(usageId string) (Discount, error) {
	args := r.Called(usageId)
	return args.Get(0).(Discount), args.Error(1)
}

func (r *MockedRepository) Save(discount Discount) error {
	args := r.Called(discount)
	return args.Error(0)
}

func (r *MockedRepository) Delete(ddiscount Discount) error {
	args := r.Called(ddiscount)
	return args.Error(0)
}

func getTestCart(id string) *cart.Cart {
	line1 := &cart.Line{
		Product: product.SpecializedProduct{
			Product:    product.Product{Id: "1"},
			Price:      100,
			Attributes: make(map[string]string),
		},
		Num:                1,
		PriceAfterDiscount: 100,
	}

	line2 := &cart.Line{
		Product: product.SpecializedProduct{
			Product:    product.Product{Id: "2"},
			Price:      200,
			Attributes: make(map[string]string),
		},
		Num:                1,
		PriceAfterDiscount: 200,
	}

	line3 := &cart.Line{
		Product: product.SpecializedProduct{
			Product:    product.Product{Id: "3"},
			Price:      300,
			Attributes: make(map[string]string),
		},
		Num:                2,
		PriceAfterDiscount: 600,
	}

	lines := []*cart.Line{line1, line2, line3}

	return &cart.Cart{
		Id:    id,
		Lines: lines,
	}
}

type MockedLockStore struct {
	mock.Mock
}

func (l *MockedLockStore) Lock(key string) *sync.RWMutex {
	args := l.Called(key)
	return args.Get(0).(*sync.RWMutex)
}

type MockedOrderService struct {
	mock.Mock
}

func (s *MockedOrderService) GetUserOrders(user *user.User) ([]order.Order, error) {
	args := s.Called(user)
	return args.Get(0).([]order.Order), args.Error(1)
}

func (s *MockedOrderService) GetUserOrdersInStatus(user *user.User, orderStatus []order.OrderStatus) ([]order.Order, error) {
	args := s.Called(user, orderStatus)
	return args.Get(0).([]order.Order), args.Error(1)
}
