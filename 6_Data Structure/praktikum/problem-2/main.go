package main

import (
	"fmt"
	"strconv"
)

// Fungsi mengecek string dan mengembalikan nilai unique
func munculSekali(angka string) []int {
	var result[]int
	count := make(map[string]int)
	
	for _, value := range angka {
		value := string(value)
		if value != " " {
			count[value]++
		}
	}

	for key, value := range count {
		if value == 1 {
			check, _ := strconv.Atoi(key)
			result = append(result, check)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}