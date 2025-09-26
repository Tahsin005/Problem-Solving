package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for T := 0; T < t; T++ {
		var n int
		fmt.Scan(&n)
		count := 0
		countm1 := 0
		for i := 0; i < n; i++ {
			temp := 0
			fmt.Scan(&temp)
			if temp == 0 {
				count++
			}
			if temp == -1 {
				countm1++
			}
		}
		if countm1 % 2 != 0 {
			count += 2
		}
		fmt.Println(count)
	}
}
