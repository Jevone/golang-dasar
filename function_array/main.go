package main

import "fmt"

func main() {
	number := [5]int{10, 4, 2, 3, 20}

	fmt.Println("Jumlah Elemen : ", len(number))
	fmt.Println("Menampilkan angka:", number[4])
	number[4] = 100
	fmt.Println("Menampilkan ulang angka", number[4])

	fmt.Println("Ini adalah array :")

	for i := 0; i < len(number); i++ {
		fmt.Println("isi index ke : ", i)
	}
}
