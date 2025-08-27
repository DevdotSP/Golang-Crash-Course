package sampleInterface

import (
	"fmt"
)

// SampleInterface serves as the entry point to run interface examples
func SampleInterface() {
	// Uncomment the samples you want to run
	// FirstSample()   // Basic interface usage
	// SecondSample()  // Multiple product types implementing the interface
	// ThirdSample()   // Includes optional methods like discount/offer
	LastExample() // Advanced example using embedded interfaces
}

// -------------------- FIRST SAMPLE --------------------
// Demonstrates simple usage of a single product type implementing the catalog interface
func FirstSample() {
	tshirt := configurable{}
	tshirt.price = 250
	tshirt.qty = 2

	fmt.Println("Shipping Charge: ", tshirt.shipping()) // Calls the catalog method
	fmt.Println("Tax: ", tshirt.tax())                 // Calls the catalog method
}

// -------------------- SECOND SAMPLE --------------------
// Demonstrates using multiple product types (configurable, simple, download)
// Each product implements catalog interface
func SecondSample() {
	// Configurable product
	tshirt := configurable{}
	tshirt.price = 250
	tshirt.qty = 2
	fmt.Println("Configurable Product")
	fmt.Println("Shipping Charge: ", tshirt.shipping())
	fmt.Println("Tax: ", tshirt.tax())

	// Simple physical product
	mobile := simple{"Samsung S-7", 10, 25}
	fmt.Println("\nSimple Product")
	fmt.Println("Shipping Charge: ", mobile.shipping())
	fmt.Println("Tax: ", mobile.tax())

	// Downloadable product (no shipping)
	book := download{"Python in 24 Hours", 19, 1}
	fmt.Println("\nDownloadable Product")
	fmt.Println("Tax: ", book.tax())
}

// -------------------- THIRD SAMPLE --------------------
// Demonstrates additional methods outside the interface (like offer/discount)
func ThirdSample() {
	// Configurable product with discount
	tshirt := configurable{}
	tshirt.price = 250
	tshirt.qty = 2
	fmt.Println("Configurable Product")
	fmt.Println("Shipping Charge: ", tshirt.shipping())
	fmt.Println("Tax: ", tshirt.tax())
	fmt.Println("Discount: ", tshirt.offer()) // Optional method, not part of the interface

	// Simple product with discount
	mobile := simple{"Samsung S-7", 3000, 2}
	fmt.Println("\nSimple Product")
	fmt.Println(mobile.name)
	fmt.Println("Shipping Charge: ", mobile.shipping())
	fmt.Println("Tax: ", mobile.tax())
	fmt.Println("Discount: ", mobile.offer())

	// Downloadable product
	book := download{"Python in 24 Hours", 50, 1}
	fmt.Println("\nDownloadable Product")
	fmt.Println(book.name)
	fmt.Println("Tax: ", book.tax())
}
