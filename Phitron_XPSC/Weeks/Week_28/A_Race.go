package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, x, y int
		fmt.Scan(&a, &x, &y)
		if a <= min(x, y) || a >= max(x, y) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
