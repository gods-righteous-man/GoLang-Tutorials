//store data in key value pair

package main

import (
	"fmt"
)

func main() {

	cart := make(map[string]int)

	//Write
	cart["milk"] = 2
	cart["eggs"] = 12
	cart["bread"] = 1
	cart["chicken"] += 1

	fmt.Println(cart)

	//Read
	chickenQuantity, check := cart["chicken"]

	fmt.Println(chickenQuantity, check)

	testQuantity, check := cart["fish"]

	fmt.Println(testQuantity, check) //check is to check whether the key is present in the map, it returns a boolean value

	//Delete

	delete(cart, "bread")

	fmt.Println(cart)
}
