package main

import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++{
		var a, b int
		fmt.Scan(&a,&b)
		if a == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(2 * b + a + 1)
		}
	}
}
