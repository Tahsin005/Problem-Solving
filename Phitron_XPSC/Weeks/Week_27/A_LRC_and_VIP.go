package main

import (
	"fmt"
	"sort"
)

func main () {
	var t int
	fmt.Scan(&t)

	for range t {
		var n int
		fmt.Scan(&n)

		numbers := make([]int, n)
		for i := range n {
			fmt.Scan(&numbers[i])
		}

		copyNumbers := make([]int, n)
		copy(copyNumbers, numbers)

		sort.Ints(copyNumbers)

		if copyNumbers[0] == copyNumbers[n - 1] {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")

			for i := range n {
				if numbers[i] == copyNumbers[n - 1] {
					fmt.Print("2")
				} else {
					fmt.Print("1")
				}

				if i != n - 1 {
					fmt.Print(" ")
				} else {
					fmt.Print("\n")
				}
			}
		}
	}
}