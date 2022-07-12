package main

import "fmt"

func main() {
	animal := "cat"

	switch animal {
	case "cat":
		{
			fmt.Println("Meow")
		}

	case "dog":
		{
			fmt.Println("whoof")

		}

	default:
		{
			fmt.Println("Default case")
		}
	}
}
