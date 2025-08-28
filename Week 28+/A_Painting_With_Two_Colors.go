package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for range t {
		var n, a, b int
		fmt.Scan(&n, &a, &b)

		var ok bool
		if a >= b && a % 2 == n % 2 && a % 2 == b % 2 {
			ok = true
		}
		if b >= a && b % 2 == n % 2 {
			ok = true
		}

		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
