package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for i := 0; i < len(numbers); i++ {
		if *numbers[i] > max {
			temp := max
			max = *numbers[i]
			*numbers[i] = temp
		}
		
		if *numbers[i] < min {
			temp := min
			min = *numbers[i]
			*numbers[i] = temp
		}
	}

	return min, max
}

// Problem 4 - Min and Max Using Pointer
func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min", min)
	fmt.Println("Nilai max", max)
}