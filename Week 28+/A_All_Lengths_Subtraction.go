package main

import "fmt"

func sol() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	l := 0
	r := n - 1
	for i := 1; i <= n; i++ {
		if p[l] == i {
			l++
		} else if p[r] == i {
			r--
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func main() {
	var tt int
	fmt.Scan(&tt)
	for i := 0; i < tt; i++ {
		sol()
	}

}
