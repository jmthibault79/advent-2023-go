package main

import "fmt"

func main() {
	a, b, c := localPointersAreFine()
	fmt.Println(a, *a, b, *b, c, *c)

	*b = 0
	fmt.Println(a, *a, b, *b, c, *c)
}

func localPointersAreFine() (a, b, c *int) {
	a = new(int)
	*a = 9
	myArray := [4]int{1, 2, 3, 4}
	b = &myArray[3]
	var bob int = 15
	c = &bob
	return
}
