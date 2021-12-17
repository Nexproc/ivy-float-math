package ivymath

import (
	"ivy_float_math/config"
	"ivy_float_math/exec"
	"ivy_float_math/value"
	"math/big"
)

var (
	constInit = false
)

type BigMath interface {
	Logn(*big.Float) *big.Float
}

type math struct {
	ctx value.Context
}

func (m math) Logn(float *big.Float) *big.Float {
	return value.FloatLog(m.ctx, float)
}

// Create new module for doing math on values from math/big
func New() BigMath {
	var c = exec.NewContext(config.NewConfig())
	if !constInit {
		value.Consts(c)
		constInit = true
	}
	return math{ctx: c}
}
