package sampleInterface

import "fmt"

// -------------------- INTERFACE HIERARCHY --------------------

// Anotherdiscount defines a method to calculate a discount/offer
type Anotherdiscount interface {
	Anotheroffer() float64
}

// Anothergiftpack defines a method to check gift pack availability
type Anothergiftpack interface {
	Anotheravailable() string
}

// Anothercatalog composes multiple interfaces and adds shipping & tax calculation
// This shows Go’s interface embedding (interface inheritance)
type Anothercatalog interface {
	Anotherdiscount
	Anothergiftpack
	Anothershipping() float64
	Anothertax() float64
}

// -------------------- STRUCT TYPES --------------------

// Anotherconfigurable represents a configurable product
// Implements Anothercatalog via all embedded methods
type Anotherconfigurable struct {
	name       string
	price, qty float64
}

// Tax for configurable product: 5%
func (c *Anotherconfigurable) Anothertax() float64 {
	return c.price * c.qty * 0.05
}

// Shipping for configurable product: $5 per quantity
func (c *Anotherconfigurable) Anothershipping() float64 {
	return c.qty * 5
}

// Offer/discount for configurable product: 15% of price
func (c *Anotherconfigurable) Anotheroffer() float64 {
	return c.price * 0.15
}

// Gift pack availability based on price
func (c *Anotherconfigurable) Anotheravailable() string {
	if c.price > 1000 {
		return "Gift Pack Available"
	}
	return "Gift Pack not Available"
}

// Anotherdownload represents downloadable products (digital goods)
type Anotherdownload struct {
	name       string
	price, qty float64
}

// Tax for downloadable product: 10%
func (d *Anotherdownload) Anothertax() float64 {
	return d.price * d.qty * 0.10
}

// Gift pack availability for downloadable products
func (d *Anotherdownload) Anotheravailable() string {
	if d.price > 500 {
		return "Gift Pack Available"
	}
	return "Gift Pack not Available"
}

// Anothersimple represents a standard physical product
type Anothersimple struct {
	name       string
	price, qty float64
}

// Tax for simple product: 3%
func (s *Anothersimple) Anothertax() float64 {
	return s.price * s.qty * 0.03
}

// Shipping for simple product: $3 per quantity
func (s *Anothersimple) Anothershipping() float64 {
	return s.qty * 3
}

// Offer/discount for simple product: 10%
func (s *Anothersimple) Anotheroffer() float64 {
	return s.price * 0.10
}

// -------------------- DEMO FUNCTION --------------------

// LastExample demonstrates using all product types and their methods
func LastExample() {
	// Configurable Product Example
	tshirt := Anotherconfigurable{}
	tshirt.price = 1550
	tshirt.qty = 2
	fmt.Println("Configurable Product")
	fmt.Println("Shipping Charge: ", tshirt.Anothershipping())
	fmt.Println("Tax: ", tshirt.Anothertax())
	fmt.Println("Discount: ", tshirt.Anotheroffer())
	fmt.Println(tshirt.Anotheravailable())

	// Simple Product Example
	mobile := Anothersimple{"Samsung S-7", 3000, 2}
	fmt.Println("\nSimple Product")
	fmt.Println("Name:", mobile.name)
	fmt.Println("Shipping Charge: ", mobile.Anothershipping())
	fmt.Println("Tax: ", mobile.Anothertax())
	fmt.Println("Discount: ", mobile.Anotheroffer())

	// Downloadable Product Example
	book := Anotherdownload{"Python in 24 Hours", 50, 1}
	fmt.Println("\nDownloadable Product")
	fmt.Println("Name:", book.name)
	fmt.Println("Tax: ", book.Anothertax())
	fmt.Println(book.Anotheravailable())
}
