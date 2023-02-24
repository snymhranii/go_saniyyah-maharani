package main

import "fmt"

func munculSekali(angka string) []int {
	count := make(map[int]int)

	for _, c := range angka {
		num := int(c - '0')
		count[num]++
	}

	result := []int{}
	for num, freq := range count {
		if freq == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
