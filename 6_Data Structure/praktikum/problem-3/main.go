package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var result[]int
	temp := make(map[int]int)
	
	for i, val := range arr {
		
		// cek apakah ada val dan memastikan ada pair sum
		if j, check := temp[val]; check {
			result = append(result, j, i)
		}

		temp[target - val] = i
	}

	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}