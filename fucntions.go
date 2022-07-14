package main

import "fmt"

func cow(num int) {
	for i := num; i < 10; i++ {
		fmt.Println(i)
	}
}

func addition(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	fmt.Println("Calling cow funcition")

	cow(1)

	fmt.Println("Calling addition function")

	sum := addition(10, 20)

	fmt.Println(sum)
}
