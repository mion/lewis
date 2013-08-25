package main

import (
	"fmt"
	. "lewis"
	"math/big"
	r "reflect"
)

// func Analyze(a interface{}) {
// 	fmt.Println("------")

// 	t := r.TypeOf(a)
// 	v := r.ValueOf(a)
// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())
// 	fmt.Println("Size:", t.Size())
// 	fmt.Println("Methods:", t.NumMethod())
// 	for i := 0; i < t.NumMethod(); i++ {
// 		fmt.Println(t.Method(i))
// 	}
// }

func convertTypeIfNecessary(t r.Type, a Any) r.Value {
	if !r.TypeOf(a).AssignableTo(t) {
		if r.TypeOf(a).ConvertibleTo(t) {
			Debug("Converting", a, "with type", r.TypeOf(a), "to type", t, "result", r.ValueOf(a).Convert(t).Interface())
			return r.ValueOf(a).Convert(t)
		} else {
			Error(r.TypeOf(a), " is not assignable nor convertible to ", t)
			return r.ValueOf(nil)
		}
	} else {
		Debug("No convertion needed for argument", a)
		return r.ValueOf(a)
	}
}

type Number struct {
	val int
}

func (n Number) Int64() int64 {
	return int64(n.val)
}

func main() {
	// does not work
	fn := r.ValueOf(big.NewInt)

	x1 := int64(5)
	x2 := int64(5)
	x3 := 5

	in1 := []r.Value{r.ValueOf(x1)}
	res1 := fn.Call(in1)[0]
	fmt.Println(res1.Interface())

	in2 := []r.Value{r.ValueOf(x2).Convert(r.TypeOf(x2))}
	// fmt.Println(in2[0].Interface(), in2[0].String())
	res2 := fn.Call(in2)[0]
	fmt.Println(res2.Interface())

	in3 := []r.Value{r.ValueOf(x3).Convert(r.TypeOf(x1))}
	res3 := fn.Call(in3)[0]
	fmt.Println(res3.Interface())

	// works
	b := fn.Call([]r.Value{r.ValueOf(int64(5))})[0]
	fmt.Println(b.Interface())

	var x int
	var y int64
	var x2y r.Value
	var y2x r.Value
	// a := 42
	// b := 3.4
	// c := "Hello"
	// // d := big.NewInt(0)
	// Analyze(a)
	// Analyze(b)
	// Analyze(c)
	// // Analyze(d)
	x = 5
	xt := r.TypeOf(x)
	yt := r.TypeOf(y)
	xv := r.ValueOf(x)
	yv := r.ValueOf(y)
	x2y = r.ValueOf(x).Convert(yt)
	y2x = r.ValueOf(y).Convert(xt)
	fmt.Println("x", x, xt, xv)
	fmt.Println("y", y, yt, yv)
	fmt.Println("x asgn to y?", xt.AssignableTo(yt))
	fmt.Println("y asgn to x?", yt.AssignableTo(xt))
	fmt.Println("x conv to y?", xt.ConvertibleTo(yt))
	fmt.Println("y conv to x?", yt.ConvertibleTo(xt))
	fmt.Println("x's value converted to y", x2y.Interface())
	fmt.Println("y's value converted to x", y2x.Interface())
}
