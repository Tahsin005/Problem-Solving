package main

import "fmt"

func main() {
	k := "qwertyuiopasdfghjkl;zxcvbnm,./"
	var keyboard [128]int

	for i := 0; i < len(k); i++ {
		keyboard[int(k[i])] = i
	}

	var d, s, r string
	fmt.Scanln(&d)
	fmt.Scanln(&s)

	for i := 0; i < len(s); i++ {
		if d == "R" {
			index := keyboard[int(s[i])]
			r += string(k[index - 1])
		} else {
			r += string(k[keyboard[int(s[i])] + 1])
		}
	}

	fmt.Println(r)
}
