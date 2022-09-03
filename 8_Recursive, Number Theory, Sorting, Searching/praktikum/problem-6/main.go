package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	if money < 1 {
		fmt.Println(0)
	} else {
		var n = len(productPrice)

		for i := 0; i < n; i++ {
			var min_index = i

			for j := i; j < n; j++ {
				if productPrice[j] < productPrice[min_index] {
					min_index = j
				}
			}

			productPrice[i], productPrice[min_index] = productPrice[min_index], productPrice[i]
		}
		count := 0

		if productPrice[0] < money {
			count++
		}

		total := productPrice[0]

		for p := 1; p < len(productPrice); p++ {
			total += productPrice[p]
			
			if total <= money {
				count++
			} else {
				break
			}
		}

		fmt.Println(count)
	}
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}