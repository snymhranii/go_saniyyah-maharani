package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 2
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	fmt.Println(ans)
}
