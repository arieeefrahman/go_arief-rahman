package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string = ""
	runes := []rune(s.name)
	var stringToRuneAscii []int
	var runeToStringAscii []string
	var totalOfAlphabet int = 26

	// Konversi string ke bentuk rune agar bisa dikonversi ke ascii.
	for i := 0; i < len(runes); i++ {
		stringToRuneAscii = append(stringToRuneAscii, int(runes[i]))
	}

	// Shifting huruf
	// Asumsi kita hanya konversi tidak lebih dari satu kata dan tanpa spasi
	for i := 0; i < len(stringToRuneAscii); i++ {
		ascii_A, ascii_Z := int('A'), int('Z')
		ascii_a, ascii_z := int('a'), int('z')
		
		// Handle uppercase
		if stringToRuneAscii[i] >= ascii_A && stringToRuneAscii[i] <= ascii_Z {
			position := stringToRuneAscii[i] - ascii_A
			position = totalOfAlphabet - position - 1
			stringToRuneAscii[i] = position + ascii_A
		}
		
		// Handle lowercase
		if stringToRuneAscii[i] >= ascii_a && stringToRuneAscii[i] <= ascii_z {
			position := stringToRuneAscii[i] - ascii_a
			position = totalOfAlphabet - position - 1
			stringToRuneAscii[i] = position + ascii_a
		}

		// Konversi rune kembali ke bentuk string
		character := string(rune(stringToRuneAscii[i]))
		runeToStringAscii = append(runeToStringAscii, character)
	}

	nameEncode = strings.Join(runeToStringAscii, "")

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode string = ""
	s.name = s.nameEncode

	// Karena pada problem ini menggunakan substitution cypher (atbash), maka enkripsi dan dekripsi akan menghasilkan nilai yang sama.
	// Maka cukup kita panggil method Encode().
	nameDecode += s.Encode()

	return nameDecode
}

// Problem 6 - Substitution Chiper
func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode  Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}