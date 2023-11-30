package main

import "fmt"

func main() {
	// arrays have value semantics (copies, not references) and fixed length
	anArray := [...]int{1, 2, 3, 4}
	fmt.Println(anArray, len(anArray))
	copiedArray := anArray
	fmt.Println(anArray == copiedArray) // true
	copiedArray[2] = 100
	fmt.Println(anArray == copiedArray) // false

	var fixedArray [3]int
	fixedArray = [3]int{7, 8, 9}
	fmt.Println(fixedArray, len(fixedArray))

	// slices have reference semantics and dynamic length,
	// and are *backed by arrays* but I don't know the implications of this yet
	aSlice := []int{1, 2, 3}
	fmt.Println(aSlice, len(aSlice))
	copiedSlice := aSlice
	//	fmt.Println(aSlice == copiedSlice) // no!  can't compare slices like this
	fmt.Println(aSlice[0] == copiedSlice[0]) // true
	copiedSlice[0] = 99
	fmt.Println(aSlice[0] == copiedSlice[0]) // true

	// other ways to define slices
	anotherSlice := make([]int, 2)
	fmt.Println(anotherSlice, len(anotherSlice))
	var emptySlice []int
	fmt.Println(emptySlice, len(emptySlice))

	notEmptySlice := append(emptySlice, 1)
	fmt.Println(notEmptySlice, len(notEmptySlice))
}
