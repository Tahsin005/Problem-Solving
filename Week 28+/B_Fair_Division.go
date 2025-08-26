package main

import "fmt"

func main() {
	var n, m, num int
	fmt.Scan(&n)
	var one int = 0
	for n != 0 {
		n--
		one = 0
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			fmt.Scan(&num)
			if num == 1 {
				one++
			}
		}
		if (m & 1 == 1 && one != 0 && one & 1 == 0) || (m & 1 == 0 && one & 1 == 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
