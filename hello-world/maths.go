package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	a = 1
	var b = 2
	c := 3
	var d = a + b + c
	fmt.Println("d is", d)

	e, f := twoDivisions(b, c)
	fmt.Println("e is", e)
	fmt.Println("f is", f)

	const g int32 = 5
	fmt.Println("g is", g)

	nan := math.NaN()
	fmt.Println(nan, math.IsNaN(nan))

	inf := math.Inf(+1)
	negInf := math.Inf(-1)
	fmt.Println(inf, negInf, math.IsInf(inf, +1), math.IsInf(negInf, +1),
		math.IsInf(inf, -1), math.IsInf(negInf, -1))
}

func twoDivisions(a, b int) (c int, d float32) {
	return a / b, float32(a) / float32(b)
}
