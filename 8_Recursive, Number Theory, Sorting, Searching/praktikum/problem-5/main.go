package main

import "fmt"

func FindMinAndMax(arr []int) string {
	if len(arr) < 1 {
		return "Tidak dapat mencari karena array kosong"
	}
	
	var min int = 0
	var max int = 0
	var index_min int = 0
	var index_max int = 0
	
	for i, v := range arr {
		if v < min {
			min = v
			index_min = i
		} else  if v > max {
			max = v
			index_max = i
		}
	}

	str := fmt.Sprintf("min: %v \tindex: %v\tmax: %v  \tindex: %v", min, index_min, max, index_max)
	
	return str
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}