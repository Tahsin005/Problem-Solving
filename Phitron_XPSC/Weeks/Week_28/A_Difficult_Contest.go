package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s []byte
		fmt.Scan(&s)
		index := 0
		for j := 0; j < len(s); j++ {
			if s[j] == 'T' {
				s[index], s[j] = s[j], s[index]
				index++
			}
		}
		fmt.Println(string(s))
	}
}
