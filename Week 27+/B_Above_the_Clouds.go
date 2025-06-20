package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for test := 0; test < t; test++ {
		var n int
		var s string
		fmt.Scan(&n, &s)

		freq := make([]int, 26)
		for _, c := range s {
			freq[c - 'a']++
		}

		found := false
		for i := 1; i <= n - 2; i++ {
			if freq[s[i] - 'a'] >= 2 {
				found = true
				break
			}
		}

		if found {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
