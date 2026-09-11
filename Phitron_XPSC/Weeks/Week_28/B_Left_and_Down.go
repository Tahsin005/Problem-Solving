package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for range n {
		var x, y, k int
		fmt.Scan(&x, &y, &k)
		g := gcd(x, y)
		x /= g
		y /= g
		switch {
		case x <= k && y <= k:
			fmt.Println(1)
		case x == y:
			fmt.Println(1)
		default:
			fmt.Println(2)
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}
