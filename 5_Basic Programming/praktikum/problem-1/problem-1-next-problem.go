package main

import "fmt"

func main() {
	r := 4.0
	var T float64

	fmt.Scanf("%g", &T)

	luas := 2 * 3.14 * r * (r + T)
	fmt.Println(luas)
}