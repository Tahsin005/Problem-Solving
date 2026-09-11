package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var str string
		fmt.Scan(&str)
		l := len(str)

		var count0, count1 int
		for j := 0; j < l; j++ {
			if str[j] == '0' {
				count0++
			} else {
				count1++
			}
		}

		min := count0
		if count1 < count0 {
			min = count1
		}

		if min & 1 == 1 {
			fmt.Println("DA")
		} else {
			fmt.Println("NET")
		}
	}
}
