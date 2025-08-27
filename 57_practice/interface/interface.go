package sampleInterface

// -------------------- INTERFACE DEFINITION --------------------

// catalog defines a contract for all types of products in the catalog.
// Any type implementing this interface must provide `shipping` and `tax` methods.
type catalog interface {
	shipping() float64
	tax() float64
}

// -------------------- STRUCT TYPES --------------------

// configurable represents a product that can be configured (like configurable electronics).
// Implements the catalog interface and has an additional offer method.
type configurable struct {
	name       string
	price, qty float64
}

// tax calculates 5% tax for configurable products
func (c *configurable) tax() float64 {
	return c.price * c.qty * 0.05
}

// shipping calculates shipping cost for configurable products
// Here, shipping is $5 per quantity
func (c *configurable) shipping() float64 {
	return c.qty * 5
}

// offer calculates discount/offer price for configurable product
func (c *configurable) offer() float64 {
	return c.price * 0.15
}

// -------------------- DOWNLOADABLE PRODUCT --------------------

// download represents a digital/downloadable product.
// Implements the catalog interface but typically has no shipping cost (could be ignored if needed)
type download struct {
	name       string
	price, qty float64
}

// tax calculates 7% tax for downloadable products
func (d *download) tax() float64 {
	return d.price * d.qty * 0.07
}

// -------------------- SIMPLE PRODUCT --------------------

// simple represents a simple physical product
// Implements the catalog interface and also has an offer method
type simple struct {
	name       string
	price, qty float64
}

// tax calculates 3% tax for simple products
func (s *simple) tax() float64 {
	return s.price * s.qty * 0.03
}

// shipping calculates shipping cost for simple products ($3 per quantity)
func (s *simple) shipping() float64 {
	return s.qty * 3
}

// offer calculates discount/offer price for simple product
func (s *simple) offer() float64 {
	return s.price * 0.10
}
