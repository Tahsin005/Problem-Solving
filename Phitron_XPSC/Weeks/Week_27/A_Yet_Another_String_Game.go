package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)

		ans := ""
		for i, c := range s {
			if i % 2 == 0 {
				if c == 'a' {
					ans += "b"
				} else {
					ans += "a"
				}
			} else {
				if c == 'z' {
					ans += "y"
				} else {
					ans += "z"
				}
			}
		}
		fmt.Println(ans)
	}
}