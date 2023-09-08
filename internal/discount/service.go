package discount

import (
	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/user"
)

// Service interface for a discounting service
type Service interface {
	Apply(*cart.Cart, string, *user.User) error
	IsEligible(*cart.Cart, string, *user.User) bool
	RollbackUsage(usageId string, discountCode string) error
	Save(Discount) error
	Delete(Discount) error
	Get(string) (Discount, error)
}
