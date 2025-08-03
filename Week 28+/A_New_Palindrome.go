package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--
		var s string
		fmt.Scan(&s)
		flag := false
		for i := 1; i < len(s) / 2; i++ {
			if s[i] != s[0] {
				flag = true
				break
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
