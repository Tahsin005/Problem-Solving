package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)
	for ii := 0; ii < tt; ii++ {
		var n int
		fmt.Scan(&n)
 
		mp := make(map[int]int)
		mx := 0
 
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			mp[x]++
			if mp[x] <= mx {
				continue;
			}
 
			mx = mp[x]
		}
 
		fmt.Println(n - mx)
	}
}