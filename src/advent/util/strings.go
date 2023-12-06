package util

import (
	"strconv"
	"strings"
)

func ParseSpacedInts(line string) (out []int, err error) {
	strVals := strings.Fields(line)
	for _, sVal := range strVals {
		iVal, err := strconv.Atoi(sVal)
		if err != nil {
			return nil, err
		}
		out = append(out, iVal)
	}
	return out, nil
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
