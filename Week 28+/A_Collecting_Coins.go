package main

import "fmt"

func main() {
	var t, a, b, c, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b, &c, &n)
		var total = a + b + c + n
		if total % 3 != 0 || total / 3 < a || total / 3 < b || total / 3 < c {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
