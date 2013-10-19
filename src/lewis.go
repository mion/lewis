package main

import (
	"bufio"
	"fmt"
	. "lewis"
	"os"
	"io/ioutil"
)

func REPL() {
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

func execute(source string) {
	scope := NewScope(GlobalScope)
	p := Parse(source)
	Eval(p, scope)
}

func main() {
	if len(os.Args) < 2 {
		REPL()	
	} else {
		filename := os.Args[1]
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("ERROR: unable to open file '", filename, "'")
		}
		source := string(contents)
		fmt.Println(source)
		execute(source)
	}
}
