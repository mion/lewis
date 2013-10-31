package main

import (
	"bufio"
	"fmt"
	. "lewis"
	"os"
	"io/ioutil"
	str "strings"
)

func startREPL() {
	fmt.Println("-- Lewis 0.1\n" +
		"-- See README.md for more information.\n" +
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

func runFile(source string) {
	scope := NewScope(GlobalScope)

	lines := str.Split(source, "\n\n")
	for _, line := range lines {
		expr := Parse(line)
		Eval(expr, scope)
	}
}

func main() {
	if len(os.Args) < 2 {
		startREPL()	
	} else {
		filename := os.Args[1]
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("[!] Error: unable to open file '", filename, "'")
		}
		source := string(data)
		runFile(source)
	}
}
