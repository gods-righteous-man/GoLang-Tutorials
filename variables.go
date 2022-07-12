package main

import "fmt"

func main() {

	fmt.Println("Hello World") //to print stuff

	var favBook = "Harry Potter"

	fmt.Println(favBook)

	var otherBook string = "Andrew Tate"

	fmt.Println(otherBook)

	var thirdBook string

	thirdBook = "Ram Dev Baba"

	var intvar int

	var boolvar bool

	var stringvar string

	fmt.Println(thirdBook)

	fmt.Println(intvar, boolvar, stringvar)

	intvar, boolvar, stringvar = 1, true, "teststring" //compound creation

	fmt.Println(intvar, boolvar, stringvar)

	var (
		one = 1
		two = "two"
	)

	fmt.Println(one, two)

	favanimal := "tiger" //not using var keyword , declare and assign

	fmt.Println(favanimal)

}
