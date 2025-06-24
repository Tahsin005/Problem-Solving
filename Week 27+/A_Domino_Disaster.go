package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Scan(&n, &s)

		var ans string
		for _, v := range s {
			if v == 'U' {
				ans += "D"
			} else if v == 'D' {
				ans += "U"
			} else if v == 'L' || v == 'R' {
				ans += string(v)
			}
		}

		fmt.Println(ans)
	}
}