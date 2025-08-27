package beginer

import "fmt"

func SampleMapstructAppend() {

	type Emp struct {
		x int
		y []string
	}
	var list = map[string]*Emp{"e1": {1001, []string{"John", "US"}}}

	e := new(Emp)
	e.x = 1002
	e.y = []string{"Rock", "UK"}

	list["e2"] = e

	fmt.Println(list["e1"])
	fmt.Println(list["e2"])

}
