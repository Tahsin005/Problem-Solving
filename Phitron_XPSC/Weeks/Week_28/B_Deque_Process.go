package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)

	for ; tt > 0; tt-- {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		i, j := 0, len(arr) - 1
		big := true
		for i < j {
			if big {
				if arr[i] > arr[j] {
					i++
					fmt.Print("L")
				} else {
					j--
					fmt.Print("R")
				}
			} else {
				if arr[i] < arr[j] {
					i++
					fmt.Print("L")
				} else {
					j--
					fmt.Print("R")
				}
			}
			big = !big
		}

		fmt.Print("L")
		fmt.Println()
	}
}