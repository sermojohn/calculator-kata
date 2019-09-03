package ops

import "sermojohn/calc-kata/logger"

type Calc struct {
	l logger.Logger
}

func NewCalc(l logger.Logger) *Calc {
	return &Calc{
		l: l,
	}
}

func (c *Calc) Add(a, b int64) int64 {
	c.l.Log("adding %d and %d", a, b)
	return a + b
}

func (c *Calc) Subtract(a, b int64) int64 {
	return a - b
}

func (c *Calc) Multiply(a, b int64) float64 {
	return float64(a) * float64(b)
}

func (c *Calc) Devide(a, b int64) float64 {
	return float64(a) / float64(b)
}
