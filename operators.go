package main

import "fmt"

func main() {

	var num1 = 1

	var num2 = 20

	var sum = num1 + num2

	fmt.Println(sum)

	var mul = num1 * num2

	var div = num1 / num2

	var rem = num2 % 10

	fmt.Println(mul, div, rem)

	//Relational

	var result = num1 < num2

	fmt.Println(result)

	//logical

	const name = "Gaurav"
	age := 23

	var check = name == "Gaurav" && age > 23

	fmt.Println(check)

}
