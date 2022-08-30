package main

import "fmt"

func cetakTabelPerkalian(number int) {
	
	//check batasan input
	if number < 30 {
		for i := 1; i <= number; i++ {
			for j := 1; j <= number; j++ {
				fmt.Printf("%d\t", i * j)
			}
			
			fmt.Println()
		}
	} else {
		fmt.Println("Input tidak boleh dari 30. Mohon input dengan benar!")
	}

	
}

func main() {
	cetakTabelPerkalian(9)
}