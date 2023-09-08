package order

import (
	"time"

	"github.com/amir-qasemi/shop-discount/internal/cart"
)

// OrderStatus the statuses an order can be in
type OrderStatus string

// Some simple statuses an order can be in
const (
	Paid      OrderStatus = "paid"
	Pending   OrderStatus = "pending"
	Processed OrderStatus = "processed"
	Shiped    OrderStatus = "shiped"
	Delivered OrderStatus = "delivered"
)

// OrderHistoryvthe lifetime of an order
type OrderHistory struct {
	Status     OrderStatus
	CreationTs time.Time
}

// Order order of a user
type Order struct {
	Username     string
	Cart         cart.Cart
	OrderHistory []OrderHistory
}
