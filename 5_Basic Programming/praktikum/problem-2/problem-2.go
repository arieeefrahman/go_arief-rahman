package main

import "fmt"

func main() {
	var studentScore int = 80
	fmt.Print("Masukkan nama mahasiswa : ")
	var nama string
	fmt.Scan(&nama)

	fmt.Println("\nNama mahasiswa : " + nama)
	if (studentScore < 0 || studentScore > 100) {
		fmt.Println("Nilai Invalid")
	} else if (studentScore >= 80) {
		fmt.Println("Nilai A")
	} else if (studentScore >= 65) {
		fmt.Println("Nilai B")
	} else if (studentScore >= 50) {
		fmt.Println("Nilai C")
	} else if (studentScore >= 35) {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Nilai E")
	}
}