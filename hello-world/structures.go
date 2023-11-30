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

	// append takes varargs, so you need to "unpack" the slice to use it here
	notEmptySlice = append(notEmptySlice, aSlice...)
	fmt.Println(notEmptySlice, len(notEmptySlice))

	notEmptySlice = append(notEmptySlice, []int{7, 8, 9}...)
	fmt.Println(notEmptySlice, len(notEmptySlice))

	// maps are dynamic and by-reference too
	map1 := map[string]int{"three": 3, "four": 4}
	map1["one"] = 1
	fmt.Println(map1, len(map1))

	map2 := map1
	map2["FOO"] = 999
	fmt.Println(map1, len(map1))
}
