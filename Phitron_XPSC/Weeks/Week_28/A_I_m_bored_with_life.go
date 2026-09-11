package main

import "fmt"

func main () {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		a = b
	}

	ans := 0
	for ans = 1; a > 0; a-- {
		ans *= a
	}

	fmt.Println(ans)
}