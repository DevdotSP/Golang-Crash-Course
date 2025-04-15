package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
return 2 * (r.width + r.height)
}

func printCicle(c circle) {
	fmt.Println("Shape:", c)
}
func printRectangle(r rectangle) {
	fmt.Println("Shape:", r)
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())
}

func main() {
	// c1 := circle{radius: 5.5}
	// r1 := rectangle{width: 3, height: 2.1}

	// printCicle(c1)
	// printRectangle(r1)

	var s shape
	fmt.Printf("%T\n",s)

	ball := circle{radius: 2.5}
	s = ball

	print(s)
	fmt.Printf("Type of s: %T\n",s)

	room := rectangle{width: 2,height: 2}
	s = room

	fmt.Printf("Type of s: %T\n",s)

}