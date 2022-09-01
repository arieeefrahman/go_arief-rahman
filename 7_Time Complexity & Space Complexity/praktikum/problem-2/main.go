package main

import (
	"fmt"
)

// function untuk menghitung nilai x^n dengan time complexity O(log n)
func pow(x, n int) int {
	result := 1
	
	// kondisi (n > 0) digunakan karena apapun yang berpangkat 0 pasti menghasilkan nilai 1 
	// asumsi bahwa input n selalu bernilai positif
	for n > 0 {
		if (n & 1) != 0 {
			result *= x
		}
		
		n /= 2
		x *= x 
 	}

	return result
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}