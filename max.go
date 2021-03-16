package main

import "fmt"

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(Max([]int{2, 1, 6, 3, 8, 32, 16}))
}
