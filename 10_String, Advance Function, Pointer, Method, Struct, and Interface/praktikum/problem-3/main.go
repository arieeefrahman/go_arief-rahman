package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Problem 3 - Swap Two Number Using Pointer
func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}