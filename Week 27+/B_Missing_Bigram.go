package main

import "fmt"

func main () {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		var s []string
		for j := 0; j < n - 2; j++ {
			var str string
			fmt.Scan(&str)
			s = append(s, str)
		}

		res := s[0]
		for i := 1; i < n - 2; i++ {
			if res[len(res) - 1] == s[i][0] {
				res += string(s[i][1])
			} else {
				res += s[i]
			}
		}
		if len(res) < n {
			res += "a"
		}

		fmt.Println(res)
	}
}