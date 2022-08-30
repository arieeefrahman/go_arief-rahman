package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Fungsi mengecek string dan mengembalikan nilai unique
func munculSekali(angka string) []int {
	count := make(map[string]int)
	chars := make([]string, 0, len(count))
	var result[]int

	for _, x := range angka {
		if string(x) != " " {
			count[string(x)]++
		}
	}

	for c := range count {
		chars = append(chars, c)
	}

    sort.Slice(chars, func(i, j int) bool {
        return count[chars[i]] > count[chars[j]]
    })

	for key, value := range count {
		check, _ := strconv.Atoi(key)

		if value == 1 {
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