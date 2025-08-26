package main

import "fmt"

func main() {
	name := "Andrei"
	fmt.Println(&name)

	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, ptr, &ptr)

	var ptr1 *float64
	_ = ptr1

	p := new(int)
	x = 100
	p = &x
	fmt.Printf("p is of type %T with value %v\n", p, p)
	fmt.Printf("address of x is %p\n", &x)

	*p = 90
	fmt.Println(*p)
	fmt.Println("*p == x:", *p == x)

	*p = 10
	*p = *p / 2
	fmt.Println(x)

	PartOne()
	PartTwo()
}

// Safe version: returning pointer created with new()
func change(a *int) *float64 {
	*a = 100
	b := new(float64)
	*b = 5.5
	return b
}

func changeVar(a int) {
	a = 66
}

func PartOne() {
	x := 8
	p := &x

	fmt.Println("value of x before calling change():", x)
	change(p)
	fmt.Println("value of x after calling change():", x)

	fmt.Println("value of x before calling changeVar():", x)
	changeVar(x)
	fmt.Println("value of x after calling changeVar():", x)
}

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.5
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["x"] = 30
}

func PartTwo() {
	quantity, price, name, sold := 5, 300.2, "Laptop", true
	fmt.Println("BEFORE changeValues():", quantity, price, name, sold)

	changeValues(quantity, price, name, sold)
	fmt.Println("AFTER changeValues():", quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER changeValuesByPointer():", quantity, price, name, sold)

	present := Product{price: 100, productName: "Watch"}

	changeProduct(present)
	fmt.Println("AFTER changeProduct():", present)

	changeProductByPointer(&present)
	fmt.Println("AFTER changeProductByPointer():", present)

	prices := []int{10, 20, 30}
	changeSlice(prices)
	fmt.Println("prices after changeSlice():", prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("myMap after changeMap():", myMap)
}
