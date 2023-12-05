package util

// inspired by roj
func MaybePanic(e error) {
	if e != nil {
		panic(e)
	}
}
