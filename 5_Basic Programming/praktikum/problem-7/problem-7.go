package main

import "fmt"

func playWithAsterisk(n int) {
	k := n - 1
	
	for i := 0; i < n; i++ {
		for j := 0;j < k; j++{
			fmt.Print(" ")	
		}
		k--

		for j := 1; j <= i + 1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterisk(5)
}