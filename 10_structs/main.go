// =============================================
// STRUCTS IN GO
// =============================================
//
// - A struct is a composite type that groups together fields.
// - Fields can be named or anonymous (field name derived from type).
// - Structs can be embedded inside other structs (composition).
// - Struct literals allow initialization by order or by field name.
// - Structs are compared by value (all fields must match).
// - Assignment creates a copy, not a reference.
//
// =============================================

package main

import (
    "fmt"
    "strings"
)

func main() {

    // =============================================
    // ANONYMOUS STRUCT
    // =============================================
    // Anonymous struct = defined on-the-fly, no named type
    diana := struct {
        firstName, lastName string
        age                 int
    }{
        firstName: "Diana",
        lastName:  "Muller",
        age:       30,
    }

    fmt.Printf("%#v\n", diana)
    // Output: struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30}

    // =============================================
    // ANONYMOUS FIELDS
    // =============================================
    // If no field names are given, the field type itself is used as the name
    type Book struct {
        string
        float64
        bool
    }

    b1 := Book{"1984 by George Orwell", 10.2, false}
    fmt.Printf("%#v\n", b1)
    // Output: main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}

    // Access anonymous field by type name
    fmt.Println(b1.string) // => 1984 by George Orwell

    // Mixing anonymous + named fields
    type Employee1 struct {
        name   string
        salary int
        bool
    }

    e := Employee1{"John", 40000, false}
    fmt.Printf("%#v\n", e)
    // => main.Employee1{name:"John", salary:40000, bool:false}

    e.bool = true // updating anonymous field

    fmt.Println(strings.Repeat("#", 10))

    // =============================================
    // EMBEDDED STRUCTS
    // =============================================
    // One struct can contain another as a field → composition
    type Contact struct {
        email, address string
        phone          int
    }

    type Employee struct {
        name        string
        salary      int
        contactInfo Contact // embedded struct field
    }

    // Struct literal with nested struct
    john := Employee{
        name:   "John Keller",
        salary: 3000,
        contactInfo: Contact{
            email:   "jkeller@company.com",
            address: "Street 20, London",
            phone:   042324234,
        },
    }

    fmt.Printf("%+v\n", john)
    // %+v prints field names with values

    // Accessing fields
    fmt.Printf("Employee's salary: %d\n", john.salary)
    fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

    // Updating nested struct field
    john.contactInfo.email = "new_email@thecompany.com"
    fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

    // Call helper function for more struct examples
    OtherStruct()
}

// =============================================
// FUNCTION WITH MORE STRUCT EXAMPLES
// =============================================
func OtherStruct() {
    // Define a named struct type
    type book struct {
        title  string
        author string
        year   int
    }

    // ---------------------------------
    // Struct literal with ordered fields
    lastBook := book{"The Divine Comedy", "Dante Aligheri", 1320}
    fmt.Println(lastBook)

    // Struct literal with named fields (order doesn’t matter)
    bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
    _ = bestBook

    // Zero values for omitted fields
    aBook := book{title: "Just a random book"}
    fmt.Printf("%#v\n", aBook) // => main.book{title:"Just a random book", author:"", year:0}

    // Access field
    fmt.Println(lastBook.title)

    // ERROR: invalid field access
    // pages := lastBook.pages // pages doesn't exist

    // Update fields
    lastBook.author = "The Best"
    lastBook.year = 2020
    fmt.Printf("lastBook: %+v\n", lastBook)

    // Compare struct values (true if all fields are equal)
    randomBook := book{title: "Random Title", author: "John Doe", year: 100}
    fmt.Println(randomBook == lastBook) // => false

    // Assignment copies struct (not reference)
    myBook := randomBook
    myBook.year = 2020 // modifying copy only
    fmt.Println(myBook, randomBook)
    // => {Random Title John Doe 2020} {Random Title John Doe 100}
}
