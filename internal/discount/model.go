package discount

import (
	"fmt"

	"github.com/amir-qasemi/shop-discount/internal/cart"
)

// Unit the type
type Unit string

const (
	Absoulute  Unit = "absoulute"
	Percentage Unit = "percentage"
)

// Discount the general structure for a discount(could be either rule based or adhoc)
type Discount interface {
	Code() string
	Unit() Unit
	MaxVal() int
}

// Usage this struct represents a usage of a discount in cart
type Usage struct {
	Id   string
	Cart *cart.Cart
}

// ErrorCode differnet error codes related to discount package
type ErrorCode int

const (
	Used              ErrorCode = 100
	RequirementNotMet ErrorCode = 101
)

// NotElligibleForDiscountErr an error for when user is not eligible for discount
type NotElligibleForDiscountErr struct {
	Reason string
	Code   ErrorCode
}

func (e *NotElligibleForDiscountErr) Error() string {
	return fmt.Sprintf("Discount is not eligible, code: %d, reason: %s", e.Code, e.Reason)
}
