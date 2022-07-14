package main

import "fmt"

func main() {

	type Animal struct {
		class  string
		age    int
		gender bool
	}

	var teddy = Animal{
		class:  "bear",
		age:    20,
		gender: true,
	}

	fmt.Println(teddy)

	teddy.age = teddy.age + 1

	fmt.Println(teddy)

	timmy := Animal{"ant", 2, false}

	fmt.Println(timmy)

}
