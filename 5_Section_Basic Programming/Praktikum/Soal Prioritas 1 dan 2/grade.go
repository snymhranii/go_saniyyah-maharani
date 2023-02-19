package main

import "fmt"

func main() {
	// Deklarasi variabel nilai
	var nilai int

	// Meminta input nilai dari user
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	// Menentukan grade berdasarkan nilai yang dimasukkan
	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Grade A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("Grade B")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("Grade C")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("Grade D")
	} else if nilai == 0 && nilai <= 34 {
		fmt.Println("Grade E")
	}
}
