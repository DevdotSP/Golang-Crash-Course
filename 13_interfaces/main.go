package main

import (
    "fmt"
    "math"
)

// -------------------------
// INTERFACE DEFINITION
// -------------------------

// shape interface defines a "contract":
// Any type that implements both `area()` and `perimeter()`
// automatically satisfies the shape interface.
type shape interface {
    area() float64
    perimeter() float64
}

// -------------------------
// STRUCT TYPES
// -------------------------

// rectangle type with width and height fields
type rectangle struct {
    width, height float64
}

// circle type with radius field
type circle struct {
    radius float64
}

// -------------------------
// METHODS FOR CIRCLE
// -------------------------

// circle implements shape's area()
func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

// circle implements shape's perimeter()
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

// circle has an extra method not in shape interface
// NOTE: `volume()` is not part of shape, so you can't
// call it directly from an interface value.
func (c circle) volume() float64 {
    return 4.0 / 3.0 * math.Pi * math.Pow(c.radius, 3)
}

// -------------------------
// METHODS FOR RECTANGLE
// -------------------------

// rectangle implements shape's area()
func (r rectangle) area() float64 {
    return r.height * r.width
}

// rectangle implements shape's perimeter()
func (r rectangle) perimeter() float64 {
    return 2 * (r.height + r.width)
}

// -------------------------
// FUNCTIONS USING INTERFACE
// -------------------------

// print accepts a shape (polymorphism).
// Works with any type that satisfies the shape interface.
func print(s shape) {
    fmt.Printf("Shape: %#v\n", s)             // %#v shows struct details
    fmt.Printf("Area: %v\n", s.area())        // method call via interface
    fmt.Printf("Perimeter: %v\n", s.perimeter())
}

// -------------------------
// MAIN FUNCTION
// -------------------------

func main() {

    // Assign a circle to an interface variable (shape).
    // `circle` implements `shape`, so this works.
    var s shape = circle{radius: 2.5}

    fmt.Printf("%T\n", s) // -> main.circle (dynamic type inside interface)

    // ✅ Only methods in the shape interface are accessible:
    fmt.Printf("Circle Area: %v\n", s.area())

    // ❌ Can't access circle-specific methods directly:
    // s.volume() // ERROR: not defined in shape

    // -------------------------
    // TYPE ASSERTIONS
    // -------------------------

    // Extract the dynamic value from the interface (downcasting).
    fmt.Printf("Sphere Volume: %v\n", s.(circle).volume())

    // Safe type assertion: check if `s` holds a circle
    ball, ok := s.(circle)
    if ok {
        fmt.Printf("Ball Volume: %v\n", ball.volume())
    }

    // -------------------------
    // TYPE SWITCH
    // -------------------------

    // A cleaner way than multiple type assertions.
    switch value := s.(type) {
    case circle:
        fmt.Printf("%#v has circle type\n", value)
    case rectangle:
        fmt.Printf("%#v has rectangle type\n", value)
    }

    // -------------------------
    // POLYMORPHISM DEMO
    // -------------------------
    fmt.Println("\nUsing the print() function with different types:")

    print(circle{radius: 3})
    print(rectangle{width: 4, height: 6})
}
