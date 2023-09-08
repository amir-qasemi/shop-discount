package discount

import (
	"github.com/amir-qasemi/shop-discount/internal/product"
)

type NewUserDiscount struct {
	generalAdhocDiscount
}

type ProductDiscount struct {
	generalAdhocDiscount
	Product product.Product
}

type MinProductDiscount struct {
	generalAdhocDiscount
	Product   product.Product
	MinNumber int
}

func (d *NewUserDiscount) Apply(req discountReqWrapper) error {
	return d.checkAndApply(req, true)
}

func (d *NewUserDiscount) IsEligible(req discountReqWrapper) bool {
	return d.checkAndApply(req, false) == nil
}

func (d *NewUserDiscount) checkAndApply(req discountReqWrapper, apply bool) error {
	if err := d.validateReq(req); err != nil {
		return err
	}

	orders, err := req.orderService.GetUserOrders(req.user)

	if err != nil {
		return err
	}

	if orders != nil && len(orders) > 0 {
		return &NotElligibleForDiscountErr{Code: RequirementNotMet, Reason: "Have orders already"}
	}

	if apply {
		for _, l := range req.cart.Lines {
			if err := d.applyToLine(l); err != nil {
				return err
			}

		}
		d.addUsage(req.cart)
	}

	return nil
}

func (d *ProductDiscount) Apply(req discountReqWrapper) error {
	return d.checkAndApply(req, true)
}

func (d *ProductDiscount) IsEligible(req discountReqWrapper) bool {
	return d.checkAndApply(req, false) == nil
}

func (d *ProductDiscount) checkAndApply(req discountReqWrapper, apply bool) error {
	if err := d.validateReq(req); err != nil {
		return err
	}

	applied := false
	for _, l := range req.cart.Lines {
		if l.Product.Id == d.Product.Id {
			if apply {
				d.applyToLine(l)
			}

			applied = true
		}
	}
	if !applied {
		return &NotElligibleForDiscountErr{Code: RequirementNotMet, Reason: "No such product exists in cart"}
	} else {
		if applied && apply {
			d.addUsage(req.cart)
		}
		return nil
	}
}

func (d *MinProductDiscount) Apply(req discountReqWrapper) error {
	return d.checkAndApply(req, true)
}

func (d *MinProductDiscount) IsEligible(req discountReqWrapper) bool {
	return d.checkAndApply(req, false) == nil
}

func (d *MinProductDiscount) checkAndApply(req discountReqWrapper, apply bool) error {
	if err := d.validateReq(req); err != nil {
		return err
	}

	numOfProds := 0

	for _, l := range req.cart.Lines {
		if l.Product.Id == d.Product.Id {
			numOfProds += l.Num
		}
	}

	if numOfProds >= d.MinNumber {
		if apply {
			for _, l := range req.cart.Lines {
				if l.Product.Id == d.Product.Id {
					d.applyToLine(l)
				}
			}
			d.addUsage(req.cart)
		}
		return nil
	} else {
		return &NotElligibleForDiscountErr{Code: RequirementNotMet, Reason: "Not enough number of products exists in cart"}
	}

}
