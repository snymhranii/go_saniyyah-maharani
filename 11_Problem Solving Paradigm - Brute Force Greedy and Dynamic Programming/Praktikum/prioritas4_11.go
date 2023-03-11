package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	// Memeriksa apakah nilai a, b, dan c valid
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Invalid input values")
		return
	}

	// Memeriksa apakah solusi ada
	if b > a {
		fmt.Println("no solution")
		return
	}

	// Menghitung nilai z
	z := int(math.Sqrt(float64(c - a*a/b)))

	// Memeriksa apakah nilai z valid
	if z*z != c-a*a/b {
		fmt.Println("1, 2, 3")
		return
	}

	// Menghitung nilai x dan y
	y := (a*z - b) / (z - 1)
	x := a - y - z

	// Menampilkan hasil
	fmt.Println(x, y, z)
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
