package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
