package main

import (
	"fmt"
	"os"
)

func main() {
	str1 := "I'm a string" // in utf-8
	fmt.Println(str1, len(str1))

	asBytes := []byte(str1)
	fmt.Println(asBytes, len(asBytes))

	str2 := `I'm a string with 
a newline`
	fmt.Println(str2, len(str2))

	// runes are unicode characters, aliased to int32
	rune1 := 'æ'
	rune2 := '\n'
	fmt.Println(rune1, rune2) // outputs the unicode number, not the character itself

	filePtr, err := os.Create("output.txt")
	fmt.Println("file pointer", filePtr)
	fmt.Println("error", err)
	followed := *filePtr
	fmt.Println("followed file", followed)

	bytesWritten, err := fmt.Fprint(filePtr, "Joel wrote this")
	fmt.Println("bytes written", bytesWritten)
	fmt.Println("error", err)

	err = filePtr.Close()
	fmt.Println("error", err)
}
