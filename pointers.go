package main

import (
	"fmt"
)

func main() {

	var a = 1

	var b = a

	b += 1

	var c = 1

	var d *int //d is a pointer and its gonna point to an integer

	d = &c

	*d += 1 // *d holds the value of c

	c = 3

	fmt.Println(a, b, c, *d)

	//instead of passing data to a function which creates a copy, create and send a pointer to the actually existing data, this saves space.
	//make sure that no changes happen from the function to the data else it will change the data
}
