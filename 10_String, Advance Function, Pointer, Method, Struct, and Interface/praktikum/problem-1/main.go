package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	str := a
	subStr := b

	// menukar nilai str dan subStr jika panjang subStr lebih panjang dari panjang str
	if len(subStr) > len(str) {
		str, subStr = subStr, str
	}

	// mengecek apakah subStr terdapat dalam str
	check := strings.Contains(str, subStr)

	output := func() string{
		if check {
			return subStr
		} else {
			return fmt.Sprintf("'%s' not contains in '%s'", subStr, str)
		}
	}

	return output()
}

// Problem 1 - Compare String
func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
