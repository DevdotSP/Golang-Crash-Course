package diffuc

import "fmt"

func LoopingFor() {
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value of index %v is %v", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}
	fmt.Println(names)

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	for range "Hello" {
		fmt.Println("Hello")
	}

}
