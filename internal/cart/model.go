package cart

import "github.com/amir-qasemi/shop-discount/internal/product"

// Cart the basic attributes needed in a cart for applying discount.
// This could be a structure defined in protocol buffers and retreived using gRPC
type Cart struct {
	// Id of the cart. Should be unqiue for each cart
	Id string
	// Each cart several lines corresponding to the added products for user
	Lines []*Line
}

// Cart the basic attributes needed in a cart for applying discount.
// This could be a structure defined in protocol buffers and retreived using gRPC
type Line struct {
	// The product of this line
	Product product.SpecializedProduct
	// Number of products added to this line. Should be greater than 1
	Num int
	// The price after the discount is applied. Initially set to Product.Price * Num
	PriceAfterDiscount int
}
