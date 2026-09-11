package main

import "fmt"

func main() {
	var s []byte
	var cnts [256]int
	fmt.Scanln(&s)
	for _, v := range s {
		cnts[v]++
	}
	oddCnts := 0
	for _, v := range cnts {
		if v % 2 != 0 {
			oddCnts++
		}
	}
	ans := "Second"
	if oddCnts == 0 || oddCnts % 2 != 0 {
		ans = "First"
	}
	fmt.Println(ans)
}
