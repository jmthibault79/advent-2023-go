package util

import (
	"strconv"
	"strings"
)

func ParseSpacedInts(str string) (out []int, err error) {
	strVals := strings.Fields(str)
	for _, sVal := range strVals {
		if sVal != " " {
			iVal, err := strconv.Atoi(sVal)
			if err != nil {
				return nil, err
			}
			out = append(out, iVal)
		}
	}
	return out, nil
}

func ParseSeparatedInts(str, sep string) (out []int, err error) {
	return ParseSpacedInts(strings.ReplaceAll(str, sep, " "))
}

func IntsToString(in []int) string {
	var b strings.Builder
	b.WriteString("[ ")
	for _, val := range in {
		b.WriteString(strconv.Itoa(val))
		b.WriteString(" ")
	}
	b.WriteString("]")
	return b.String()
}
