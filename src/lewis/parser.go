package lewis

import (
	"container/list"
	"regexp"
	scv "strconv"
	"strings"
)

func Categorize(s string) Any {
	if i, err := scv.ParseInt(s, 0, 64); err == nil {
		return i
	} else if str, err := scv.Unquote(s); err == nil {
		return str
	} else {
		return NewSymbol(s)
	}
}

func Tokenize(s string) *list.List {
	var leftParens = regexp.MustCompile(`\(`)
	var rightParens = regexp.MustCompile(`\)`)

	s = leftParens.ReplaceAllString(s, " ( ")
	s = rightParens.ReplaceAllString(s, " ) ")
	s = strings.TrimSpace(s)

	l := list.New()
	for _, str := range strings.Fields(s) {
		l.PushBack(str)
	}
	return l
}

func Parenthesize(tokens *list.List) Any {
	if tokens.Len() == 0 {
		return nil
	}
	t := tokens.Front().Value.(string)
	tokens.Remove(tokens.Front())
	if t == "(" {
		c := Cons(nil, nil)
		p := c
		for tokens.Front().Value.(string) != ")" {
			p.Car = Parenthesize(tokens)
			p.Cdr = Cons(nil, nil)
			p = p.Cdr
		}
		tokens.Remove(tokens.Front())
		p.Cdr = nil
		return c
	} else if t == ")" {
		panic("unexpected )")
	} else {
		x := Categorize(t)
		return x
	}
}

func Parse(input string) Any {
	return Parenthesize(Tokenize(input))
}
