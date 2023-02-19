package main

import "fmt"

func main() {
	var bilangan int
	var nama string

	fmt.Printf("masukan nama : ")
	fmt.Scanf("%s\n", &nama)
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&bilangan)

	fmt.Printf("hai %v\nFaktor dari %v\nadalah ", nama, bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
