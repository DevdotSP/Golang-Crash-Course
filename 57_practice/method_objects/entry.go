package methodobjects

import "fmt"

// -------------------- ENUM-LIKE TYPE --------------------

// gadgets represents a feature set of a smartphone
type gadgets uint8

// Predefined gadget constants using iota
const (
	Camera gadgets = iota
	Bluetooth
	Media
	Storage
	VideoCalling
	Multitasking
	Messaging
)

// -------------------- STRUCT DEFINITIONS --------------------

// mobile represents a generic phone with basic info
type mobile struct {
	make  string
	model string
}

// smartphone embeds gadgets to represent advanced phones
type smartphone struct {
	gadgets gadgets
}

// launch is a method with receiver type *smartphone
func (s *smartphone) launch() {
	fmt.Println("New Smartphone Launched:")
}

// android struct embeds mobile + smartphone and adds waterproof info
type android struct {
	mobile
	smartphone
	waterproof string
}

func (a *android) samsung() {
	fmt.Printf("%s %s\n", a.make, a.model)
}

// iphone struct embeds mobile + smartphone and adds sensor info
type iphone struct {
	mobile
	smartphone
	sensor int
}

func (i *iphone) apple() {
	fmt.Printf("%s %s\n", i.make, i.model)
}

// ObjectsSample demonstrates struct embedding + method usage
func ObjectsSample() {
	t := &android{}
	t.make = "Samsung"
	t.model = "Galaxy J7 Prime"
	// Adding all gadgets together (bitwise-like behavior)
	t.gadgets = Camera + Bluetooth + Media + Storage + VideoCalling + Multitasking + Messaging
	t.launch()
	t.samsung()
}

// -------------------- NON-STRUCTURE TYPE METHOD --------------------

// multiply is a custom type with methods
type multiply int

// tentimes demonstrates a method on a non-struct type
func (m multiply) tentimes() int {
	return int(m * 10)
}

func NonStructureTypeMethod() {
	var num int
	fmt.Print("Enter any positive integer: ")
	fmt.Scanln(&num)
	mul := multiply(num)
	fmt.Println("Ten times of a given number is: ", mul.tentimes())
}

// -------------------- METHOD SET DEMO --------------------

// salary → total → hra → basic
// These demonstrate method chaining between different custom types

type salary float64
type total float64
type hra float64
type basic float64

func (s salary) total() total {
	return total(s)
}

func (t total) hra() hra {
	t += t * 0.3 // 30% HRA addition
	return hra(t)
}
func (t total) salary() salary {
	t -= t * 0.10 // 10% Tax deduction
	return salary(t)
}

func (h hra) basic() basic {
	h += h * 0.3 // 30% further addition
	return basic(h)
}
func (h hra) total() total {
	return total(h)
}

func (b basic) total() total {
	return total(b)
}

// Demonstrates chained salary calculations
func StructureTypeSet() {
	fmt.Println("Salary calculation for First Employee:")
	sal1 := basic(9000.00)
	fmt.Println(sal1.total())
	fmt.Println(sal1.total().hra().total())
	fmt.Println(sal1.total().hra().total().salary())

	fmt.Println("\nSalary calculation for Second Employee:")
	sal2 := basic(5000.00)
	fmt.Println(sal2.total())
	fmt.Println(sal2.total().salary())
}

// -------------------- METHOD "OVERLOADING" --------------------

// Go doesn’t support real method overloading,
// but you can achieve polymorphism with different types having same method name.

type Thirdmultiply int
type addition int

func (m *Thirdmultiply) twice() {
	*m = Thirdmultiply(*m * 2)
}

func (a *addition) twice() {
	*a = addition(*a + *a)
}

func MethodOverloading() {
	var mul Thirdmultiply = 15
	mul.twice()
	fmt.Println(mul)

	var add addition = 15
	add.twice()
	fmt.Println(add)
}

// Another set of types with same method name "twice"
type Anothermultiply int
type Anotheraddition int

func (m *Anothermultiply) twice() {
	*m = Anothermultiply(*m * 2)
}

func (a *Anotheraddition) twice() {
	*a = Anotheraddition(*a + *a)
}

// Standalone main-like demo for these types
func main() {
	var mul Anothermultiply = 15
	mul.twice()
	fmt.Println(mul)

	var add Anotheraddition = 15
	add.twice()
	fmt.Println(add)
}

// -------------------- ENTRY FUNCTION --------------------

// EntryObjectMethod acts as a single entry point to run demonstrations
func EntryObjectMethod() {
	// NonStructureTypeMethod()
	// StructureTypeSet()
	ObjectsSample()
}
