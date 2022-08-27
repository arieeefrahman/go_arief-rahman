package main

import "fmt"

func primeNumber(number int) bool {
	check := true
	
	if number > 1 {
		for i := 2; i <= number / 2; i++ {
			if (number % i == 0) {
				check = false
				break
			}
		}
	}

	return check
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}