package main

import "fmt"

func palindrome (input string) bool {
	check := false
	var temp string
	
	for i:= len(input) - 1; i >= 0; i-- {
		temp = temp + string(input[i])
	}

	if (temp == input) {
		check = true
	}

	return check
}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}