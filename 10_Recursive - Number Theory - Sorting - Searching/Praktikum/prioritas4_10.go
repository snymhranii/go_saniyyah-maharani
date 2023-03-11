package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	if cards[0][0] == deck[0] || cards[0][1] == deck[0] {
		return deck
	} else if cards[len(cards)-1][0] == deck[0] || cards[len(cards)-1][1] == deck[0] {
		return []int{deck[1], deck[0]}
	} else if cards[0][0] == deck[1] || cards[0][1] == deck[1] {
		return deck
	} else if cards[len(cards)-1][0] == deck[1] || cards[len(cards)-1][1] == deck[1] {
		return deck
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 4}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
