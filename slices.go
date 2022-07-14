package main

import "fmt"

func main() {

	purchases := [5]float32{1, 2, 3, 4, 5}

	mySlice := purchases[:]

	mySlice = append(mySlice, 10, 9, 8, 7)

	fmt.Println(mySlice)

	otherSlice := purchases[1:]

	fmt.Println(otherSlice)

	otherSlice = append(otherSlice, mySlice...) //equivalent to (otherSlice, 1,2,3,4)

	fmt.Println(otherSlice)
}
