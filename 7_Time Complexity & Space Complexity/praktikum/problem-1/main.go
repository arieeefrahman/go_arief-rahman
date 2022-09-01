package main

import (
	"fmt"
	"math"
)

// fungsi untuk mengecek apakah suatu bilangan termasuk bilangan prima atau bukan dengan time complexity O(sqrt(n))
func primeNumber(number int) bool {
	check := true
	
	// bilangan yang kurang dari 2 bukanlah bilangan prima
	if number < 2 {
		check = false
	} else {
		for i := 2; i <= int(math.Sqrt(float64(number)) + 1); i++ {

			// cek prima
			if (number % i == 0) {
				check = false
				break
			}
		}
	}

	return check
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}