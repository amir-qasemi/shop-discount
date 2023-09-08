// Discount
package discount

import (
	"errors"
	"fmt"
	"time"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/user"
	"github.com/amir-qasemi/shop-discount/internal/util"
)

// AdHocDiscount the interface which should be implemeted by each new discounting policy
type AdHocDiscount interface {
	Discount
	Apply(discountReqWrapper) error
	IsEligible(discountReqWrapper) bool
	Usages() map[string]Usage
}

// discountReqWrapper a wrapper to easily add more fields needed for cheking and applying discount without changing all the policies.
type discountReqWrapper struct {
	cart         *cart.Cart
	user         *user.User
	orderService order.Service
}

// generalAdhocDiscount a struct that contains field shared by all of the discounting policies.
// This struct can be embeded in new policies.
type generalAdhocDiscount struct {
	DiscountCode string
	CreationTs   time.Time
	ValidNum     int // Number of time this discount can be used. -1 for infinite usage
	XUsages      map[string]Usage
	XValue       int // The value of the discount. Should be interpreted alongside XUnit. If XUnit is percent, should be between 0 to 100.
	XUnit        Unit
	XMaxVal      int // The maximum absolute value a iscount can have
}

func (d *generalAdhocDiscount) Value() int {
	return d.XValue
}

func (d *generalAdhocDiscount) Unit() Unit {
	return d.XUnit
}

func (d *generalAdhocDiscount) MaxVal() int {
	return d.XMaxVal
}

func (d *generalAdhocDiscount) addUsage(cart *cart.Cart) {
	id := util.RandStringRunes(10)
	d.XUsages[id] = Usage{Cart: cart, Id: id}
}

func (d *generalAdhocDiscount) canUse() bool {
	return (len(d.XUsages) < d.ValidNum) || d.ValidNum == -1
}

// validateReq validates whether this discount can be used or not
func (d *generalAdhocDiscount) validateReq(req discountReqWrapper) error {
	if len(req.cart.Lines) < 1 {
		return errors.New("Discount cannot be applied on empty cart")
	}

	if !d.canUse() {
		return &NotElligibleForDiscountErr{Code: Used, Reason: "Discount is used before"}
	}

	return nil
}

// applyToLine reduce the appropriate amount from the final price.
func (d *generalAdhocDiscount) applyToLine(l *cart.Line) error {
	if d.XUnit == Absoulute {
		l.PriceAfterDiscount = l.Product.Price*l.Num - d.Value()
	} else if d.XUnit == Percentage {
		afterDiscount := int(float64(l.Product.Price*l.Num) * (float64((100 - d.Value())) / 100))
		if afterDiscount > d.XMaxVal {
			l.PriceAfterDiscount = l.PriceAfterDiscount - d.XMaxVal
		} else {
			l.PriceAfterDiscount = afterDiscount
		}

	} else {
		return errors.New(fmt.Sprintf("Discount is not applicable: Unkown unit %s", d.XUnit))
	}

	if l.PriceAfterDiscount < 0 {
		l.PriceAfterDiscount = 0
	}

	return nil
}

func (d *generalAdhocDiscount) Code() string {
	return d.DiscountCode
}

func (d *generalAdhocDiscount) Usages() map[string]Usage {
	return d.XUsages
}
