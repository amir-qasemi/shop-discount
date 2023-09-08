package discount

import (
	"time"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/product"
)

type DiscountTestCase struct {
	discountUnderTest AdHocDiscount
	prevOrders        []order.Order
	cart              *cart.Cart
	linePrices        []int
	eligibility       bool
	applyErr          bool
	usageNum          int
}

func GetTestCases() map[string]DiscountTestCase {
	// Test cases: these are just a number of sample test cases and it's not an exhaustive list
	testCases := map[string]DiscountTestCase{
		"NewUserDiscount - without any previous order - unlimited usage - without previous usage - absolute unit": {
			discountUnderTest: &NewUserDiscount{
				generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     -1,
					XUsages:      make(map[string]Usage),
					XValue:       200,
					XUnit:        Absoulute,
					XMaxVal:      -1,
				},
			},
			prevOrders:  []order.Order{},
			cart:        getTestCart("1"),
			linePrices:  []int{0, 0, 400},
			eligibility: true,
			applyErr:    false,
			usageNum:    1,
		}, "NewUserDiscount - without any previous order - overused discount - absolute unit": {
			discountUnderTest: &NewUserDiscount{
				generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages: map[string]Usage{
						"1": {Id: "1"},
					},
					XValue:  200,
					XUnit:   Absoulute,
					XMaxVal: -1,
				},
			},
			prevOrders:  []order.Order{},
			cart:        getTestCart("1"),
			linePrices:  []int{100, 200, 600},
			eligibility: false,
			applyErr:    true,
			usageNum:    1,
		}, "NewUserDiscount - without any previous order - unlimited usage - without previous usage - precentage unit": {
			discountUnderTest: &NewUserDiscount{
				generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages:      make(map[string]Usage),
					XValue:       40,
					XUnit:        Percentage,
					XMaxVal:      200,
				},
			},
			prevOrders:  []order.Order{},
			cart:        getTestCart("1"),
			linePrices:  []int{60, 120, 400},
			eligibility: true,
			applyErr:    false,
			usageNum:    1,
		},
		"NewUserDiscount - with previous order - unlimited usage - without previous usage": {
			discountUnderTest: &NewUserDiscount{
				generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages:      make(map[string]Usage),
					XValue:       200,
					XUnit:        Absoulute,
					XMaxVal:      -1,
				},
			},
			prevOrders:  []order.Order{{Username: "test"}},
			cart:        getTestCart("1"),
			linePrices:  []int{100, 200, 600},
			eligibility: false,
			applyErr:    true,
			usageNum:    0,
		},
		"NewUserDiscount- on empty cart - unlimited usage - without previous usage": {
			discountUnderTest: &NewUserDiscount{
				generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages:      make(map[string]Usage),
					XValue:       200,
					XUnit:        Absoulute,
					XMaxVal:      -1,
				},
			},
			prevOrders:  []order.Order{},
			cart:        &cart.Cart{Id: "1", Lines: []*cart.Line{}},
			linePrices:  []int{},
			eligibility: false,
			applyErr:    true,
			usageNum:    0,
		},
		"ProductDiscount - with product in cart": {
			discountUnderTest: &ProductDiscount{
				generalAdhocDiscount: generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages:      make(map[string]Usage),
					XValue:       80,
					XUnit:        Absoulute,
					XMaxVal:      -1,
				},
				Product: product.Product{Id: "1"},
			},
			prevOrders:  []order.Order{},
			cart:        getTestCart("1"),
			linePrices:  []int{20, 200, 600},
			eligibility: true,
			applyErr:    false,
			usageNum:    1,
		},
		"MinProductDiscount - with product in cart - min requirment met": {
			discountUnderTest: &MinProductDiscount{
				generalAdhocDiscount: generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages:      make(map[string]Usage),
					XValue:       80,
					XUnit:        Absoulute,
					XMaxVal:      -1,
				},
				Product:   product.Product{Id: "3"},
				MinNumber: 2,
			},
			prevOrders:  []order.Order{},
			cart:        getTestCart("1"),
			linePrices:  []int{100, 200, 520},
			eligibility: true,
			applyErr:    false,
			usageNum:    1,
		},
		"MinProductDiscount - with product in cart - min requirment not met": {
			discountUnderTest: &MinProductDiscount{
				generalAdhocDiscount: generalAdhocDiscount{
					DiscountCode: "123",
					CreationTs:   time.Now(),
					ValidNum:     1,
					XUsages:      make(map[string]Usage),
					XValue:       80,
					XUnit:        Absoulute,
					XMaxVal:      -1,
				},
				Product:   product.Product{Id: "1"},
				MinNumber: 2,
			},
			prevOrders:  []order.Order{},
			cart:        getTestCart("1"),
			linePrices:  []int{100, 200, 600},
			eligibility: false,
			applyErr:    true,
			usageNum:    0,
		},
	}
	return testCases
}
