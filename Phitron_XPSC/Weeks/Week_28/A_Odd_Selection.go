package main

import "fmt"

func solve() {
	var n, x int
	fmt.Scan(&n, &x)

	even, odd := 0, 0
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num % 2 == 0 {
			even++
		} else {
			odd++
		}
	}

	for i := 1; i <= min(odd, x); i += 2 {
		need := x - i
		if need <= even {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
}
