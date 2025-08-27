package diffuc

import "fmt"

// ExampleInterface demonstrates various ways to use interfaces in Go
func ExampleInterface() {

	// ----------------------------
	// Example 1: Simple interface with employee
	// ----------------------------
	var e1 InterfaceEmployee1
	e1 = Emp(1) // Emp satisfies InterfaceEmployee1

	// Call method with string parameter
	e1.PrintName("John Doe")

	// Call method with return value
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))

	// ----------------------------
	// Example 2: Type assertion and interface satisfaction
	// ----------------------------
	var p Polygons = Pentagon(50) // Pentagon satisfies Polygons
	p.Perimeter()                 // Call Perimeter via interface

	// Type assertion to access Pentagon-specific methods
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()

	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()

	// ----------------------------
	// Example 3: Interfaces with common methods
	// ----------------------------
	var bmw Vehicle
	bmw = Car("World Top Brand")

	var labour Human
	labour = Man("Software Developer")

	// Match parts of Vehicle and Human
	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}

	// ----------------------------
	// Example 4: Pointer receivers and interface
	// ----------------------------
	var b Book                                 // Declare Book instance
	var m Magazine                             // Declare Magazine instance
	b.Assign("Jack Rabbit", "Book of Rabbits") // Assign values to Book
	m.Assign("Rabbit Weekly", 26)              // Assign values to Magazine

	var i Printer // Declare interface variable
	fmt.Println("Call interface:")

	i = &b // Book has pointer receiver, interface works with pointer
	i.Print()

	i = &m // Magazine also satisfies Printer interface
	i.Print()

	// ----------------------------
	// Example 5: Polymorphism with multiple shapes
	// ----------------------------
	compute() // Calculate edges * multiplier
	another() // Demonstrates interface embedding
}

// compute demonstrates polymorphism using Geometry interface
func compute() {
	p := new(DifferentPentagon)
	h := new(DifferentHexagon)
	o := new(DifferentOctagon)
	d := new(DifferentDecagon)

	shapes := [...]Geometry{p, h, o, d}

	for _, shape := range shapes {
		fmt.Println(Parameter(shape, 5))
	}
}

// another demonstrates interface embedding
func another() {
	p := new(AnotherPentagon)
	h := new(AnotherHexagon)
	o := new(AnotherOctagon)
	d := new(AnotherDecagon)

	polygons := [...]AnotherPolygons{p, h, o, d}
	for _, poly := range polygons {
		fmt.Println(poly.Edges())
	}
}

// ----------------------------
// Geometry interface and structs
// ----------------------------
type Geometry interface {
	Edges() int // Returns number of edges
}

type DifferentPentagon struct{}
type DifferentHexagon struct{}
type DifferentOctagon struct{}
type DifferentDecagon struct{}

func (p DifferentPentagon) Edges() int { return 5 }
func (h DifferentHexagon) Edges() int  { return 6 }
func (o DifferentOctagon) Edges() int  { return 8 }
func (d DifferentDecagon) Edges() int  { return 10 }

// Parameter calculates perimeter/size based on edges
func Parameter(geo Geometry, value int) int {
	return geo.Edges() * value
}

// ----------------------------
// Pointer receiver example with Book and Magazine
// ----------------------------
type Book struct {
	author, title string
}

type Magazine struct {
	title string
	issue int
}

func (b *Book) Assign(author, title string) {
	b.author = author
	b.title = title
}
func (b *Book) Print() {
	fmt.Printf("Author: %s, Title: %s\n", b.author, b.title)
}

func (m *Magazine) Assign(title string, issue int) {
	m.title = title
	m.issue = issue
}
func (m *Magazine) Print() {
	fmt.Printf("Title: %s, Issue: %d\n", m.title, m.issue)
}

// Printer interface
type Printer interface {
	Print()
}

// ----------------------------
// Employee interface examples
// ----------------------------
type InterfaceEmployee interface {
	PrintName() string
	PrintAddress(id int)
	PrintSalary(basic int, tax int) float64
}

type InterfaceEmployee1 interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	// Returns salary after tax deduction
	return basic - ((basic * tax) / 100)
}

// ----------------------------
// Multiple interfaces example with polygons
// ----------------------------
type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon:", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

// ----------------------------
// Common Method Interfaces
// ----------------------------
type Vehicle interface {
	Structure() []string
	Speed() string
}

type Human interface {
	Structure() []string
	Performance() string
}

type Car string

func (c Car) Structure() []string {
	return []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Tank"}
}
func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Man string

func (m Man) Structure() []string {
	return []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
}
func (m Man) Performance() string {
	return "8 Hrs/Day"
}

// ----------------------------
// Interface embedding example
// ----------------------------
type AnotherGeometry interface {
	Edges() int
}

type AnotherPolygons interface {
	AnotherGeometry
}

type AnotherPentagon int
type AnotherHexagon int
type AnotherOctagon int
type AnotherDecagon int

func (p AnotherPentagon) Edges() int { return 5 }
func (h AnotherHexagon) Edges() int  { return 6 }
func (o AnotherOctagon) Edges() int  { return 8 }
func (d AnotherDecagon) Edges() int  { return 10 }
