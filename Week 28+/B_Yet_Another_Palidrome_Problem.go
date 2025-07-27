package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)

	for T := 0; T < t; T++ {
		fmt.Scan(&n)
		a := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		m := make(map[int]int)
		q := 0
		for i := 0; i < n; i++ {
			val, ok := m[a[i]]
			if ok == true {
				if (i - val) > 1 {
					q = 1
					break
				}
			} else {
				m[a[i]] = i
			}
		}

		if q == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
