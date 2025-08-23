package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Print(s)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}
