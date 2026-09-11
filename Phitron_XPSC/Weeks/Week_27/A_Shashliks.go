package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var k, a, b, x, y int
		fmt.Scan(&k, &a, &b, &x, &y)
		if x > y {
			a, b = b, a
			x, y = y, x
		}
		bbqA := (k - a) / x + 1
		if k < a {
			bbqA = 0
		}
		leftTemp := k - bbqA * x
		bbqB := (leftTemp-b) / y + 1
		if leftTemp < b {
			bbqB = 0
		}
		fmt.Println(bbqA + bbqB)
	}
}
