package main

import "fmt"

func main() {

	age := 34

	if age > 23 {
		fmt.Println("You are allowed")
	} else if age > 20 {
		fmt.Println("Almost reached")

	} else {
		fmt.Println("not allowed")
	}
}
