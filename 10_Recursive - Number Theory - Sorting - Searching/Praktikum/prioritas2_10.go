package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}
	pairs := make([]pair, len(counts))
	i := 0
	for name, count := range counts {
		pairs[i] = pair{name, count}
		i++
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
	return pairs
}

func main() {
	items1 := []string{"je", "js", "golang", "ruby", "ruby", "js", "js"}
	items2 := []string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}
	items3 := []string{"football", "basketball", "tenis"}

	pairs1 := MostAppearItem(items1)
	for _, pair := range pairs1 {
		fmt.Printf("%v->%v ", pair.name, pair.count)
	}
	fmt.Println()

	pairs2 := MostAppearItem(items2)
	for _, pair := range pairs2 {
		fmt.Printf("%v->%v ", pair.name, pair.count)
	}
	fmt.Println()

	pairs3 := MostAppearItem(items3)
	for _, pair := range pairs3 {
		fmt.Printf("%v->%v ", pair.name, pair.count)
	}
	fmt.Println()
}
