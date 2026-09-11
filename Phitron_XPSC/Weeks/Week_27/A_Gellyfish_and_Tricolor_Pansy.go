package main

import (
	"fmt"
)

func main() {
	var t int

	fmt.Scan(&t)
	for tc := 0; tc < t; tc++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)

		if (a < b && a < d) || (c < b && c < d) {
			fmt.Print("Flower\n")
		} else {
			fmt.Print("Gellyfish\n")
		}
	}
}
