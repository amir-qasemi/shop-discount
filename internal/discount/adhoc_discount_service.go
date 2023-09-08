package discount

import (
	"errors"
	"fmt"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/lock"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/user"
)

// AdHocDiscountService a implemntation of discount service which uses a adhoc approach(see README.md)
// The logic for evaluting whether a discount can be applied to a cart is embeded in the AdhocDiscount itself
type AdHocDiscountService struct {
	DiscountRepository Repository     // persistance layer for discounts
	OrderService       order.Service  // a service used for retrieving orders of user
	LockStore          lock.LockStore // a lock manager. If a TXN is choosed(with serializable isolation level), this store can be ommited.
}

func (s *AdHocDiscountService) Apply(cart *cart.Cart, discountCode string, user *user.User) error {
	s.LockStore.Lock(discountLockKey(discountCode))
	defer s.LockStore.Unlock(discountLockKey(discountCode))

	discount, err := s.DiscountRepository.GetDiscountByCode(discountCode)
	if err != nil {
		return err
	}

	addhocDiscount := discount.(AdHocDiscount)

	if err := addhocDiscount.Apply(discountReqWrapper{cart: cart, orderService: s.OrderService, user: user}); err != nil {
		return err
	}

	err = s.DiscountRepository.Save(discount)

	return err
}

func (s *AdHocDiscountService) IsEligible(cart *cart.Cart, discountCode string, user *user.User) bool {
	s.LockStore.Lock(discountLockKey(discountCode))
	defer s.LockStore.Unlock(discountLockKey(discountCode))

	discount, err := s.DiscountRepository.GetDiscountByCode(discountCode)
	if err != nil {
		return false
	}

	addhocDiscount := discount.(AdHocDiscount)

	return addhocDiscount.IsEligible(discountReqWrapper{cart: cart, orderService: s.OrderService, user: user})
}

func (s *AdHocDiscountService) RollbackUsage(usageId string, discountCode string) error {
	s.LockStore.Lock(discountLockKey(discountCode))
	defer s.LockStore.Unlock(discountLockKey(discountCode))

	discount, err := s.DiscountRepository.GetDiscountByUsageId(usageId)
	if discount.Code() != discountCode {
		return errors.New("Usage with the given discount code does not exist")
	}

	addhocDiscount := discount.(AdHocDiscount)
	if err != nil {
		return err
	}

	delete(addhocDiscount.Usages(), usageId)
	err = s.DiscountRepository.Save(discount)

	return err
}

func (s *AdHocDiscountService) Save(discount Discount) error {
	return s.DiscountRepository.Save(discount)
}

func (s *AdHocDiscountService) Delete(discount Discount) error {
	return s.DiscountRepository.Delete(discount)
}

func (s *AdHocDiscountService) Get(code string) (Discount, error) {
	return s.DiscountRepository.GetDiscountByCode(code)
}

func discountLockKey(code string) string {
	return fmt.Sprint("discount-%", code)
}
