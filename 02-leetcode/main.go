package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func scoreOfString(s string) int {
	res := 0
	previous_letter := 0
	for _, c := range s {
		current_letter := int(c)
		if previous_letter != 0 {
			res += abs(previous_letter - current_letter)
		}
		previous_letter = current_letter
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	res := scoreOfString(input[:len(input)-1])
	fmt.Printf("%d", res)
}
