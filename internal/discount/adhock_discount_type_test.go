package discount

import (
	"fmt"
	"testing"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Tests for different discounting policies

func TestAdhockDiscountTypes(t *testing.T) {
	mockedOrderSrv := new(MockedOrderService)

	mockCall := mockedOrderSrv.On("GetUserOrders", mock.Anything, mock.Anything).Return([]order.Order{}, nil)

	discountReqGen := func(s order.Service, cart *cart.Cart) discountReqWrapper {
		return discountReqWrapper{
			cart: cart,
			user: &user.User{
				Username: "test",
			},
			orderService: s,
		}
	}

	testCases := GetTestCases()

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			mockCall.Unset()
			mockCall = mockedOrderSrv.On("GetUserOrders", mock.Anything, mock.Anything).Return(tc.prevOrders, nil)

			discountReq := discountReqGen(mockedOrderSrv, tc.cart)

			assert.Equal(tc.eligibility, tc.discountUnderTest.IsEligible(discountReq))

			err := tc.discountUnderTest.Apply(discountReq)
			assert.Equal(tc.applyErr, err != nil, "Applying discount had error")
			assert.Equal(tc.usageNum, len(tc.discountUnderTest.Usages()), "wrong number of usages in discount")
			assert.Equal(len(tc.linePrices), len(tc.cart.Lines), "Wrong test case: number of line prices provided in test cases does not match the number of lines in the cart")
			for i, linePrice := range tc.linePrices {
				assert.Equal(linePrice, tc.cart.Lines[i].PriceAfterDiscount, fmt.Sprint("Line %i price does not match", i))
			}

			/* _, ok := tc.discountUnderTest.(*NewUserDiscount)
			if len(tc.cart.Lines) > 0 && ok {
				mockedOrderSrv.AssertExpectations(t)
			} */
		})
	}

}
