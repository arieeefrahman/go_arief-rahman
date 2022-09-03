package main

import (
	"fmt"
)

func playDomino(card [][]int, deck[]int) interface{} {
	var result[]int
	var check_biggest int = 0

	for i := 0; i < len(card); i++ {
		for j := 0; j < len(deck); j++ {
			for k := 0; k < 2; k++ {
				if card[i][j] == deck[k] {
					temp := 0
					temp_str := card[i]

					for l := 0; l < len(card[i]); l++ {
						temp+= temp_str[l]
					}

					if temp > check_biggest {
						check_biggest = temp
						result = card[i]
					}	
				}
			}
		}
	}

	if len(result) == 0 {
		return "tutup kartu"
	}

	return result
}

func main() {
	fmt.Println(playDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}