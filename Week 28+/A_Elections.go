package main

import "fmt"

func Solve() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mx := max(b, c)
	fmt.Printf("%d ", max(0, mx - a + 1))

	mx = max(a, c)
	fmt.Printf("%d ", max(0, mx - b + 1))

	mx = max(a, b)
	fmt.Printf("%d\n", max(0, mx - c + 1))
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		Solve()
		t--
	}
}
