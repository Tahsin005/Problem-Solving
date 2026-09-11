package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)

	for ; tt > 0; tt-- {
		var n int
		fmt.Scan(&n)

		minDigit := 9
		for n != 0 {
			minDigit = min(minDigit, n % 10)
			n /= 10
		}

		fmt.Println(minDigit)
	}
}