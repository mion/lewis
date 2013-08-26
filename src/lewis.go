package main

import (
	"bufio"
	"fmt"
	. "lewis"
	"os"
)

func REPL() {
	fmt.Println("-- Lewis 0.1\n" +
		"-- Type \"tutorial\" or \"example\" for more information.\n" +
		"-- Use Ctrl+D to exit.")
	scope := NewScope(GlobalScope)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		if input, err := reader.ReadString('\n'); err == nil {
			p := Parse(input[:len(input)-1])
			e := Eval(p, scope)
			fmt.Println("=>", e)
		} else {
			fmt.Println("\nREPL terminated.")
			return
		}
	}
}

func main() {
	REPL()
}
