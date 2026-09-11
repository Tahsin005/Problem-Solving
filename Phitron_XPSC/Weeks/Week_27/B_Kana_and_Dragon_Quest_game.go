package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x, n, m int
		fmt.Scan(&x, &n, &m)

		for ; x > 20 && n > 0; {
			x /= 2
			x += 10
			n -= 1
		}

		if m * 10 >= x {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}