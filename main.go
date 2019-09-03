package main

import (
	"sermojohn/calc-kata/logger"
	"sermojohn/calc-kata/ops"
)

func main() {
	lg := &logger.StdLogger{}
	c := ops.NewCalc(lg)
	lg.Log("add result %v", c.Add(1, 2))
	lg.Log("subtract result %v", c.Subtract(1, 2))
	lg.Log("multiple result %v", c.Multiply(1, 2))
	lg.Log("decide result %v", c.Devide(1, 2))
}
