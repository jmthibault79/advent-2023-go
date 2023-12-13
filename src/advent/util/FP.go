package util

func AllEqual(a []int, val int) bool {
	for _, v := range a {
		if v != val {
			return false
		}
	}
	return true
}
