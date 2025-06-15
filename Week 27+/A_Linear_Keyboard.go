package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		keyboard := scanner.Text()

		scanner.Scan()
		word := scanner.Text()

		keyPos := make(map[rune]int)
		for pos, char := range keyboard {
			keyPos[char] = pos
		}

		totalTime := 0
		for i := 1; i < len(word); i++ {
			currPos := keyPos[rune(word[i])]
			prevPos := keyPos[rune(word[i - 1])]
			totalTime += abs(currPos - prevPos)
		}

		fmt.Println(totalTime)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
