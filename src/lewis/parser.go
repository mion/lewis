package lewis

import (
	"regexp"
	scv "strconv"
	"strings"
)

func Categorize(s string) (*Atom, error) {
	if i, err := scv.ParseInt(s, 0, 64); err == nil {
		return MakeLiteral(i), nil
	} else if str, err := scv.Unquote(s); err == nil {
		return MakeLiteral(str), nil
	} else {
		return MakeIdentifier(s), nil
	}
}

func Tokenize(s string) []string {
	var leftParens = regexp.MustCompile(`\(`)
	var rightParens = regexp.MustCompile(`\)`)

	s = leftParens.ReplaceAllString(s, " ( ")
	s = rightParens.ReplaceAllString(s, " ) ")
	s = strings.TrimSpace(s)

	return strings.Fields(s)
}

// func Parenthesize(input Any, l *Cell) Any {
// 	if l == nil {
// 		return Parenthesize(input, Cons(nil, nil))
// 	} else {
// 		token := input.Shift()
// 		if token == nil {
// 			return l.Pop()
// 		} else if token == "(" {
// 			l.Push(Parenthesize(input, New()))
// 			return Parenthesize(input, l)
// 		} else if token == ")" {
// 			return l
// 		} else {
// 			str := token.(string)
// 			return Parenthesize(input, l.Concat(categorize(str)))
// 		}
// 	}
// }

func Parse(input string) Any {
	return nil
}
