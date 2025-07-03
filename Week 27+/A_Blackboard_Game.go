package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)
	for i := 0; i < tt; i++ {
		var n int
		fmt.Scan(&n)
		if n % 4 == 0 {
			fmt.Println("Bob")
		} else {
			fmt.Println("Alice")
		}
	}
}