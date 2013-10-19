package lewis

import (
	"math/big"
	"fmt"
)

type Library struct {
	functions map[string]Any
}

var BigIntLibrary = &Library{functions: map[string]Any{
	"NewInt": big.NewInt,
}}

var FmtLibrary = &Library{functions: map[string]Any {
	"Println": fmt.Println,
}}

var Libraries = map[string]*Library{
	"math/big": BigIntLibrary,
	"fmt": FmtLibrary,
}
