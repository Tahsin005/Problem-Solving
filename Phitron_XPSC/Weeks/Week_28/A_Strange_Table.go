package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var x, y, z int
		fmt.Scan(&x, &y, &z)
		var a, b, c, d int
		a = x * y
		b = a - z
		c = b % x
		d = b / x
		fmt.Println(a - ((c * y) + (d * 1)))
	}
}
