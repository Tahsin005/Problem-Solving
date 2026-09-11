package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)
	for t := 0; t < tt; t++ {
		var n, k int
		fmt.Scan(&n, &k)
		var s string
		fmt.Scan(&s)

		ans := make([]byte, k)
		mapCount := make(map[int]int)

		for _, val := range s {
			mapCount[int(val)]++
		}

		for i := 0; i < k; i++ {
			var j int
			for j = 'a'; j < 'a' + n / k && mapCount[j] > 0; j++ {
				mapCount[j]--
			}
			ans[i] = byte(j)
		}

		fmt.Println(string(ans))
	}
}