package main

import (
    "fmt"
)

// simple function that prints a message
func foo() {
    fmt.Println("This is foo()!")
}

// simple function that prints a message
func bar() {
    fmt.Println("This is bar()!")
}

// simple function that prints a message
func foobar() {
    fmt.Println("This is foobar()!")
}

func main() {

    // 🔹 Defer statement postpones the execution of a function
    // until the surrounding function returns. In this case, the surrounding
    // function is main().
    
    // defer foo() -> foo() will execute **last**, just before main() exits
    defer foo()
    
    // immediately executed function
    bar()
    
    fmt.Println("Just a string after deferring foo() and calling bar()")
    
    // defer another function
    defer foobar() // foobar() will execute before foo(), in **reverse order** of defers

    /*
       Execution order of the program:
       1. bar() is called immediately -> prints "This is bar()!"
       2. fmt.Println(...) is called -> prints "Just a string after deferring foo() and calling bar()"
       3. foobar() executes (deferred last) -> prints "This is foobar()!"
       4. foo() executes (deferred first) -> prints "This is foo()!"
       
       ✅ Key Point: Deferred functions execute **LIFO (Last In, First Out)** order.
    */
}
