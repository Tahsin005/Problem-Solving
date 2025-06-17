package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for tt := 0; tt < t; tt++ {
		var s, t string
		fmt.Scan(&s, &t)
		vec := []byte(s)
		j := 0
		for i := 0; i < len(s); i++ {
			if t[j] == s[i] {
				j++
			} else if s[i] == '?' {
				vec[i] = t[j]
				j++
			}
			if j == len(t) {
				for k := i + 1; k < len(s); k++ {
					if vec[k] == '?' {
						vec[k] = 'a'
					}
				}
				break
			}
		}
		ss := string(vec)
		if j == len(t) {
			fmt.Printf("YES\n%s\n", ss)
		} else {
			fmt.Println("NO")
		}
	}
}