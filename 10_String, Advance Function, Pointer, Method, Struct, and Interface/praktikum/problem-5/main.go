package main

import (
	"fmt"
)

type Student struct {
	name 	[]string
	score 	[]int
}

func (s Student) Average() float64 {
	var sum int = 0

	for i := 0; i < len(s.score); i++ {
		sum += s.score[i]
	}

	avg := float64(sum) / float64(len(s.score))
	
	return avg
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]

	for i := 0; i < len(s.score); i++ {
		if s.score[i] < min {
			tempMin := min
			min = s.score[i]
			s.score[i] = tempMin

			tempName := name
			name = s.name[i]
			s.name[i] = tempName
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]

	for i := 0; i < len(s.score); i++ {
		if s.score[i] > max {
			tempMax := max
			max = s.score[i]
			s.score[i] = tempMax

			tempName := name
			name = s.name[i]
			s.name[i] = tempName
		}
	}

	return max, name
}

// Problem 5 - Students Score
func main() {
	var a = Student{}

	number := 1
	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input ", number, " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input ", a.name[i], " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
		number++
	}

	fmt.Println("\n\nAverage Score Student is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student is : " + nameMax + " (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student is : " + nameMin + " (", scoreMin, ")")
}