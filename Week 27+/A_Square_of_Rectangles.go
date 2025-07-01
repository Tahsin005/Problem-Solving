package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var l1, b1, l2, b2, l3, b3 int
		fmt.Scan(&l1, &b1, &l2, &b2, &l3, &b3)

		if l1 == l2 && l1 == l3 && l1 == b1 + b2 + b3 {
			fmt.Println("YES")
		} else if b1 == b2 && b1 == b3 && b1 == l1 + l2 + l3 {
			fmt.Println("YES")
		} else if b2 == b3 && l1 == l2 + l3 && l1 == b1 + b2 {
			fmt.Println("YES")
		} else if l2 == l3 && b1 == b2 + b3 && b1 == l1 + l2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
