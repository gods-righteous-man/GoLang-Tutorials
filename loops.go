package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Hello world")
	}

	j := 1

	for j <= 10 {
		fmt.Println("Its Siraj")
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Println("Not Continued")
	}
}
