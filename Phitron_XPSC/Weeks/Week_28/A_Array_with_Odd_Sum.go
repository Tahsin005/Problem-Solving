package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)

	for ; tt != 0; tt-- {
		var n, sum int
		fmt.Scan(&n)
		var odd, even bool
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			sum += x
			odd = odd || (x % 2 != 0)
			even = even || (x % 2 == 0)
		}

		if sum % 2 != 0 || (odd && even) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}