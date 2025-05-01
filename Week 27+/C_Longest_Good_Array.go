package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var l, r int
		fmt.Scan(&l, &r)
		x, q := l, 1
		for x < r {
			x += q
			if x > r {
				break
			}
			q += 1
		}
		fmt.Println(q)
		t--
	}
}
