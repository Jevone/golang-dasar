package main

import "fmt"

func main() {
	var umur int

	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&umur)

	if umur < 13 {
		fmt.Println("Anak-Anak")
	} else if umur < 18 {
		fmt.Println("Remaja")
	} else if umur < 60 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Lansia")
	}
}
