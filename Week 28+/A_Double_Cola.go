package main

import "fmt"

func main() {
	var n int
	names := [...]string{"Sheldon", "Leonard", "Penny", "Rajesh", "Howard"}
	fmt.Scan(&n)
	p := 1
	for n > (5 * p) {
		n -= (5 * p)
		p *= 2
	}
	fmt.Println(names[(n - 1) / p])
}
