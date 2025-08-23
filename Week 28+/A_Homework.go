package main

import "fmt"

func solve() {
	var n, m int
	var s, t, i string
	fmt.Scan(&n, &s, &m, &t, &i)
	for idx, ch := range i {
		if ch == 'V' {
			s = string(t[idx]) + s
		} else {
			s = s + string(t[idx])
		}
	}
	fmt.Println(s)
}

func main() {
	var tt int
	fmt.Scan(&tt)
	for i := tt; i >= 1; i-- {
		solve()
	}
}
