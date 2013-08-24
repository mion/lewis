package lewis

import (
	"fmt"
)

var Debugging = true

func Debug(a ...interface{}) {
	if !Debugging {
		return
	}
	fmt.Println("Debugger:")
	for i, v := range a {
		if i%2 == 0 {
			fmt.Print("\t")
		} else {
			fmt.Print("\t\t")
		}
		fmt.Println(v)
	}
}

func Error(a ...interface{}) interface{} {
	args := ""
	for _, v := range a {
		args += fmt.Sprint(v)
	}
	panic("[!] Evaluation error: " + args)
	return nil
}
