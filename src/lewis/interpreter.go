package lewis

// import (
// 	"fmt"
// 	r "reflect"
// 	"strings"
// )

func Eval(x Any, s *Scope) Any {
	return nil
	// 	if sym, ok := x.(*Symbol); ok {
	// 		// variable reference
	// 		if e := env.Find(name); e != nil {
	// 			return e.Get(name)
	// 		} else {
	// 			return nil
	// 		}
	// 	}
	// 	if cell, ok := x.(*Cell); !ok {
	// 		// constant literal
	// 		return x
	// 	} else {
	// 		if cell.Car == "quote" {
	// 			return cell.Cdr.Car
	// 		} else if cell.Car == "if" {
	// 			test := cell.Cdr.Car
	// 			conseq := cell.Cdr.Car.Cdr.Car
	// 			alt := cell.Cdr.Car.Cdr.Car.Cdr.Car
	// 			// TODO: check length
	// 			if Eval(test, env) {
	// 				return Eval(conseq, env)
	// 			} else {
	// 				return Eval(alt, env)
	// 			}
	// 		} else if cell.Car == "set!" {
	// 			v := cell.Cdr.Car
	// 			exp := cell.Cdr.Car.Cdr.Car
	// 			env.Find(v)
	// 		}
	// 	}

	// if x.Kind() == r.String {
	// 	str := x.String()
	// 	if strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"") {
	// 		// string literal
	// 		return r.ValueOf(str[1 : len(str)-1])
	// 	} else {
	// 		// variable reference
	// 		if e := env.Find(str); e != nil {
	// 			return e.Get(str)
	// 		} else {
	// 			panic("undefined variable: " + str) // panic..?
	// 		}
	// 	} // else ... convert to constant literals
	//    } else if x.Kind() == r.Struct && x.Type() ==
	// } else if x.Index(0).Interface() == "quote" {
	// 	// (quote exp)
	// 	return x.Index(1)
	// } else if x.Index(0).Interface() == "if" {
	// 	// (if test conseq alt)
	// 	test := x.Index(1)
	// 	conseq := x.Index(2)
	// 	alt := x.Index(3)
	// 	// something like (= 1 2) must return a Go bool type!
	// 	if Eval(test, env).Bool() {
	// 		return Eval(conseq, env)
	// 	} else {
	// 		return Eval(alt, env)
	// 	}
	// } else if x.Index(0).Interface() == "set!" {
	// 	// (set! var exp)
	// 	v := x.Index(1)
	// 	exp := x.Index(2)
	// 	env.Find(v.Interface()).Set(v.Interface(), Eval(exp, env))
	// } else if x.Index(0).Interface() == "define" {
	// 	// (define var exp)
	// 	v := x.Index(1)
	// 	exp := x.Index(2)
	// 	env.Set(v.Interface(), Eval(exp, env))
	// } else if x.Index(0).Interface() == "lambda" {
	// 	// (lambda (var*) exp)
	// 	vars := x.Index(1)
	// 	exp := x.Index(2)
	// 	lambda := func(args []r.Value) r.Value {
	// 		Eval(exp, Env(vars, args, env))
	// 	}
	// 	return lambda
	// } else if x.Index(0).Interface() == "begin" {
	// 	// (begin exp*)
	// 	var val r.Value
	// 	for i := 1; i < x.Len(); i++ {
	// 		val = Eval(x.Index(i), env)
	// 	}
	// 	return val
	// } else {
	// 	// (proc exp*)

	// }
	// return r.ValueOf(nil)
}
