package main

import "fmt"

func main() {
	n, s := 0, ""
	fmt.Scan(&n, &s)

	result := ""
	if n % 2 == 1 {
		result, s, n = s[0:1], s[1:], n - 1
	}

	for i := 0; i < n; i += 2 {
		result = s[i:i + 1] + result + s[i + 1:i + 2]
	}
	fmt.Println(result)
}
