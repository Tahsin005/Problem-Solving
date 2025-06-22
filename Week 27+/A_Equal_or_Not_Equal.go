package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)

		cnt := 0
		for _, item := range s {
			if string(item) == "N" {
				cnt++
			}
		}

		if cnt == 1 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
