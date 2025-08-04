package main

import "fmt"

var s string

func solve() {
	fmt.Scan(&s)
	mp := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if s[i] == 'r' || s[i] == 'g' || s[i] == 'b' {
			mp[s[i]-'a'] = true
		} else {
			if mp[s[i]-'A'] == false {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
