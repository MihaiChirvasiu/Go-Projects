package main

import "fmt"

func main() {
	slice := newSlice()
	for element := range slice {
		if element%2 == 0 {
			fmt.Printf("%v is even", element)
			fmt.Println()
		} else {
			fmt.Printf("%v is odd", element)
			fmt.Println()
		}
	}
}

func newSlice() []int {
	slice := []int{}
	for i := 0; i < 11; i++ {
		slice = append(slice, i)
	}
	return slice

}
