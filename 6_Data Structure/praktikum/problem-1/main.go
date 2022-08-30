package main

import (
	"fmt"
)

// Fungsi menggabungkan 2 array dan memiliki nilai unique
func ArrayMerge(arrayA, arrayB []string) []string {
	var temp = []string{}

	temp = append(arrayA, arrayB...)
	
	check := make(map[string]bool)

	var result []string

	for _, str := range temp {
		if _, yes := check[str]; !yes {
			check[str] = true
			result = append(result, str)
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}