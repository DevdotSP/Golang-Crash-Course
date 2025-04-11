package main

import "fmt"

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python!You don't use curly braces but indentation!")
	case "Go", "golang":
		fmt.Println("Good for you, You are using curly braces{}")
	default:
		fmt.Println("Any other language is good to start")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number")
	default:
		fmt.Println("New here!")
	}

	day := 8

	switch day {
	case 1:
		fmt.Println("Today is monday")
	case 2:
		fmt.Println("Today is tuesday")
	case 3:
		fmt.Println("Today is wednesday")
	case 4:
		fmt.Println("Today is thursday")
	case 5:
		fmt.Println("Today is friday")
	case 6:
		fmt.Println("Today is saturday")
	case 7:
		fmt.Println("Today is sunday")
	default:
		fmt.Println("error")
	}
	

}
