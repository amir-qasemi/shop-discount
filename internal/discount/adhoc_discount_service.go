package discount

import (
	"errors"
	"fmt"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/lock"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/user"
)

type AdHocDiscountService struct {
	DiscountRepository Repository
	OrderService       order.Service
	LockStore          lock.LockStore
}

func (s *AdHocDiscountService) Apply(cart *cart.Cart, discountCode string, user *user.User) error {
	l := s.LockStore.Lock(discountLockKey(discountCode))
	l.Lock()
	defer l.Unlock()
	discount, err := s.DiscountRepository.GetDiscountByCode(discountCode)
	if err != nil {
		return err
	}

	addhocDiscount := discount.(AdHocDiscount)

	if err := addhocDiscount.Apply(discountReqWrapper{cart: cart, orderService: s.OrderService, user: user}); err != nil {
		return err
	}

	s.DiscountRepository.Save(discount)

	return nil
}

func (s *AdHocDiscountService) IsEligible(cart *cart.Cart, discountCode string, user *user.User) bool {
	l := s.LockStore.Lock(discountLockKey(discountCode))
	l.Lock()
	defer l.Unlock()

	discount, err := s.DiscountRepository.GetDiscountByCode(discountCode)
	if err != nil {
		return false
	}

	addhocDiscount := discount.(AdHocDiscount)

	return addhocDiscount.IsEligible(discountReqWrapper{cart: cart, orderService: s.OrderService, user: user})
}

func (s *AdHocDiscountService) RollbackUsage(usageId string, discountCode string) error {
	l := s.LockStore.Lock(discountLockKey(discountCode))
	l.Lock()
	defer l.Unlock()

	discount, err := s.DiscountRepository.GetDiscountByUsageId(usageId)
	if discount.Code() != discountCode {
		return errors.New("Usage with the given discount code does not exist")
	}

	addhocDiscount := discount.(AdHocDiscount)
	if err != nil {
		return err
	}

	delete(addhocDiscount.Usages(), usageId)
	s.DiscountRepository.Save(discount)

	return nil
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
