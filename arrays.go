package main

import "fmt"

func main() {

	purchase := [5]float32{9.9, 1.3, 4.5, 6.6}

	fmt.Println(purchase)

	temp := purchase[2]

	fmt.Println(temp)

	purchase[2] = 20.11

	fmt.Println(purchase)

	var sales [4]int

	sales[0] = 1

	fmt.Println(sales)

	testarray := [...]float32{1, 2, 3}

	fmt.Println(testarray)

	for i := 0; i < len(purchase); i++ {
		fmt.Println(purchase[i])
	}
}
