package discount

import (
	"fmt"
	"sync"
	"testing"

	"github.com/amir-qasemi/shop-discount/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// tests for AdhocDiscountService
func TestAdhocDiscountService(t *testing.T) {
	mockedRepository := new(MockedRepository)
	mockedOrderService := new(MockedOrderService)
	mockedLockStore := new(MockedLockStore)

	repMockGetDiscountByCodeCall := mockedRepository.On("GetDiscountByCode", mock.Anything).Return(nil, nil)
	mockedRepository.On("Save", mock.Anything).Return(nil)
	orderMockCall := mockedOrderService.On("GetUserOrders", mock.Anything).Return(nil, nil)

	mockedLockStore.On("Lock", mock.Anything).Return(&sync.RWMutex{})

	sut := AdHocDiscountService{
		DiscountRepository: mockedRepository,
		OrderService:       mockedOrderService,
		LockStore:          mockedLockStore,
	}

	testCases := GetTestCases()
	usr := &user.User{Username: "test"}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			repMockGetDiscountByCodeCall.Unset()
			repMockGetDiscountByCodeCall = mockedRepository.On("GetDiscountByCode", mock.Anything, mock.Anything).Return(tc.discountUnderTest, nil)

			orderMockCall.Unset()
			orderMockCall = mockedOrderService.On("GetUserOrders", usr).Return(tc.prevOrders, nil)

			assert.Equal(tc.eligibility, sut.IsEligible(tc.cart, tc.discountUnderTest.Code(), usr))

			err := sut.Apply(tc.cart, tc.discountUnderTest.Code(), &user.User{Username: "test"})
			assert.Equal(tc.applyErr, err != nil, "Applying discount had error")
			assert.Equal(tc.usageNum, len(tc.discountUnderTest.Usages()), "wrong number of usages in discount")
			assert.Equal(len(tc.linePrices), len(tc.cart.Lines), "Wrong test case: number of line prices provided in test cases does not match the number of lines in the cart")
			for i, linePrice := range tc.linePrices {
				assert.Equal(linePrice, tc.cart.Lines[i].PriceAfterDiscount, fmt.Sprint("Line %i price does not match", i))
			}

		})
	}

}
