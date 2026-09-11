package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)
	for ; tt > 0; tt-- {
		var a, b, x, y int
		fmt.Scan(&a, &b, &x, &y)

		ans := 0
		if b < a {
			if a - b == 1 {
				if a % 2 == 1 {
					ans = y
				} else {
					ans = -1
				}
			} else {
				ans = -1
			}
		} else {
			for ; a < b; a++ {
				if a % 2 == 0 {
					if y < x {
						ans += y
					} else {
						ans += x
					}
				} else {
					ans += x
				}
			}
		}
		fmt.Println(ans)
	}
}