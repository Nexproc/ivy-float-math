package ivymath

import (
	"github.com/Nexproc/ivy-float-math/config"
	"github.com/Nexproc/ivy-float-math/exec"
	"github.com/Nexproc/ivy-float-math/value"
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
