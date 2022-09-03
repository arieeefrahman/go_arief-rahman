package main

import "fmt"

func primeX(number int) int {
	ranges := 30
	var prime[]bool

	for i:=0; i <= ranges; i++ {
		prime = append(prime, true)
	}

	for j := 2; j*j <= ranges; j++ {
		if prime[j] {
			for k := j * j; k <= ranges; k += j {
				prime[k] = false
			}
		}
	}

	var temp []int

	for i, v := range prime {
		if i < 2 {
			continue
		}
		
		if v {
			temp = append(temp, i)
		}
	}

	var result int
	
	for i, v := range temp {
		if i == number - 1 {
			result = v
		}
	}

	return result
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}