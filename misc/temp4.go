package main

import (
	"fmt"
)

func shift(tokens []string) (string, []string) {
	return tokens[0], tokens[1:]
}

func main() {
	t := make([]string, 3)
	t[0] = "("
	t[1] = "x"
	t[2] = ")"
	var token string
	for len(t) > 0 {
		token, t = shift(t)
		fmt.Println(token)
	}
}
