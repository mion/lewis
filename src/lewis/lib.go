package lewis

import (
	"math/big"
)

type Library struct {
	functions map[string]Any
}

var BigIntLibrary = &Library{functions: map[string]Any{
	"NewInt": big.NewInt,
}}

var Libraries = map[string]*Library{
	"math/big": BigIntLibrary,
}
