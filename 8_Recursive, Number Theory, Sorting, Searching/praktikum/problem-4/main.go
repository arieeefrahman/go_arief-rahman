package main

import "fmt"

func MaxSequence(arr []int) int {
	var total int = 0

	if len(arr) % 2 == 0 {
		tengah := len(arr) / 2 - 1
		
		for i := tengah - 1; i <= tengah + 3; i++ {
			total += arr[i]
		}
	} else {
		tengah := len(arr) / 2
		
		for i := tengah - 1; i <= tengah + 2; i++ {
			total += arr[i]
		}
	}

	return total
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}