package main

import "fmt"

func ArrayMarge(arrayA, arrayB []string) []string {
	var ArrayMarge []string

	for _, element := range arrayA {
		ArrayMarge = append(ArrayMarge, element)
	}

	for _, element := range arrayB {
		ArrayMarge = append(ArrayMarge, element)
	}

	return ArrayMarge

}

func main() {
	fmt.Println(ArrayMarge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMarge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMarge([]string{"alisa", "yoshimitau"}, []string{"devil jin", "yoshimitau", "alisa"}))
	fmt.Println(ArrayMarge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMarge([]string{}, []string{}))

}
