package main

import (
	"fmt"
)

func Compare(a, b string) string {
	result := ""
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			result += string(a[i])
		} else {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
