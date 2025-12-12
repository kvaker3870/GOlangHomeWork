package main

import "fmt"

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if expectedOpen, isClose := pairs[char]; isClose {

			if len(stack) == 0 || stack[len(stack)-1] != expectedOpen {
				println(string(stack), string(expectedOpen))
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func main() {

	fmt.Println(isValid("({))()[]"))
}
