package main

import "fmt"

func main() {
	var bilangan int
	var nama string

	fmt.Println("=====<PROGRAM BILANGAN GANJIL GENAP>=====")
	fmt.Printf("masukan nama anda => ")
	fmt.Scanf("%s\n", &nama)
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 {
		fmt.Println("=====OUTPUT=====")
		fmt.Printf("hallo %s\n", nama)
		fmt.Printf("Bilangan %d adalah genap", bilangan)
	} else {
		fmt.Println("=====OUTPUT=====")
		fmt.Printf("hallo %s\n", nama)
		fmt.Printf("Bilangan %d adalah ganjil", bilangan)
	}
}
