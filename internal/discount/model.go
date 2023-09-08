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

type Discount interface {
	Code() string
	Unit() Unit
	MaxVal() int
}

type Usage struct {
	Id   string
	Cart *cart.Cart
}

type ErrorCode int

const (
	Used              ErrorCode = 100
	RequirementNotMet ErrorCode = 101
)

type NotElligibleForDiscountErr struct {
	Reason string
	Code   ErrorCode
}

func (e *NotElligibleForDiscountErr) Error() string {
	return fmt.Sprintf("Discount is not eligible, code: %d, reason: %s", e.Code, e.Reason)
}
