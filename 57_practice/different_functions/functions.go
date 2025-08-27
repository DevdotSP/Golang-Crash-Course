package diffuc

import "fmt"

func ExampleFunction() {
	// total, y, x := SimpleMath(5, 5)
	// add(5,5)

	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	Update(&age, &text)

	fmt.Println("After :", text, age)

	sqaureRoot := SquareSum(3)(3)(3)
	fmt.Println(sqaureRoot)
}

func add(x, y int) int {
	total := 0
	total = x + y
	return total
}

func SimpleMath(param1 int, param2 int) (returnTotal int, returnY int, returnX int) {
	returnTotal = 0
	returnY = param1
	returnX = param2
	returnTotal = (param1 * param2) + param2
	fmt.Println(returnTotal)
	fmt.Println(returnX)
	fmt.Println(returnY)

	if value := returnTotal; value == 0 {
		return 0, 0, 0
	}
	return returnTotal, param1, param2
}

// simple way to show deferencing
func Update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

type First func(int) int
type Second func(int) First

func SquareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}


