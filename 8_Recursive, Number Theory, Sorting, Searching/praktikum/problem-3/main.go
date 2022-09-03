package main

import (
	"fmt"
	// "math"
)

func primaSegiEmpat(high, wide, start int) {
	var prime[]bool
	var prime_num []int
	var result []int
	var ranges int = 1000

	for i:=0; i <= ranges; i++ {
		prime = append(prime, true)
	}

	// cek prima
	for j := 2; j*j <= ranges; j++ {
		if prime[j] {
			for k := j * j; k <= ranges; k += j {
				prime[k] = false
			}
		}
	}

	// mengambil nilai prima
	for i, v := range prime {
		if i < 1 {
			continue
		}
		
		if v {
			prime_num = append(prime_num, i)
		}
	}

	for i := 0; i < len(prime_num); i++ {
		if prime_num[i] >= start {
			result = prime_num[(i + 1): (i + 1 + (high * wide))]
			break
		}
	}

	total, x := 0, 0

	//print bentuk segi empat
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
				fmt.Printf("%3d", result[x])
				total += result[x]
				x++
				
		}

		fmt.Println()
		
	}

	fmt.Printf(" %d\n\n", total)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}