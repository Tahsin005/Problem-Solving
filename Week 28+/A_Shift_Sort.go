package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Scan(&n, &s)
		zeroCount := 0
		for _, c := range s {
			if c == '0' {
				zeroCount++
			}
		}
		onesInFirst := 0
		for i := 0; i < zeroCount; i++ {
			if s[i] == '1' {
				onesInFirst++
			}
		}
		fmt.Println(onesInFirst)
	}
}