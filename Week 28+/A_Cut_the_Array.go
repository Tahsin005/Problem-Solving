package main

import "fmt"

func solve () {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	if sum % 3 != 0 {
		fmt.Println("0 0")
		return
	}
	fmt.Printf("1 %d\n", n - 1)
}

func main () {
	var tt int
	fmt.Scan(&tt)
	for ; tt > 0; tt-- {
		solve()
	}
}