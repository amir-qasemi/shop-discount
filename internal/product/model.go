package product

// Product a struct representing products in the shop
// This could be a structure defined in protocol buffers and retreived using gRPC
type Product struct {
	// Id of the product. Must be unique
	Id string
	// Other attributes needed for discount evalution
	// Name string
}

// SpecializedProduct a product with price and attributes(e.g. color,..)
// This could be a structure defined in protocol buffers and retreived using gRPC
type SpecializedProduct struct {
	// Product base product
	Product
	// Price the price of this kind of product
	Price int
	// Attributes some application defined attributes of the prodcut(e.g. color,...)
	Attributes map[string]string
}
