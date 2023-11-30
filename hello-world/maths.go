package main

import "fmt"

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
}

func twoDivisions(a, b int) (c int, d float32) {
	return a / b, float32(a) / float32(b)
}
