package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)
	for ; tt > 0; tt-- {
		var a int
		fmt.Scan(&a)

		if (360 % (180 - a)) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}