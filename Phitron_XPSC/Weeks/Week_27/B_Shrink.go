package main

import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(1)
}
func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}