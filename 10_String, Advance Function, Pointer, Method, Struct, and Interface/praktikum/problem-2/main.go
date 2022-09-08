package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	runes := []rune(input)
	var stringToRune []int
	var runeToString []string
	var totalOfAlphabet int = 26

	// Konversi string ke bentuk rune agar bisa dikonversi ke ascii.
	for i := 0; i < len(runes); i++ {
		stringToRune = append(stringToRune, int(runes[i]))
	}

	//Shifting huruf
	for i := 0; i < len(stringToRune); i++ {
		stringToRune[i] += (offset % totalOfAlphabet)
		ascii_z := int('z')
		fmt.Println(int('A'))

		// Shifting character kembali ke a jika telah mencapai z
		if stringToRune[i] > ascii_z {
			stringToRune[i] -= totalOfAlphabet
		}

		// Konversi rune kembali ke bentuk string
		character := string(rune(stringToRune[i]))
		runeToString = append(runeToString, character)
	}
	
	result := strings.Join(runeToString, "")
	return result
}

// Problem 2 - Caesar Chiper
func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}