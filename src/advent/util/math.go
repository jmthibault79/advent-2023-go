package util

import "math"

// Greatest Common Divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor#Euclid's_algorithm
func GCD2(a, b int) int {
	if a == b {
		return a
	}

	if a < b {
		if a == 0 {
			return b
		} else {
			return GCD2(a, b%a)
		}
	} else {
		return GCD2(b, a)
	}
}

// Greatest Common Divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor#Euclid's_algorithm
func GCD(nums []int) int {
	switch len(nums) {
	case 0:
		panic("can't do GCD on empty slice")
	case 1:
		return nums[0]
	case 2:
		return GCD2(nums[0], nums[1])
	default:
		return GCD2(nums[0], GCD(nums[1:]))
	}
}

func LCM2(a, b int) int {
	return a * b / GCD2(a, b)
}

// Least Common Multiple
// https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
func LCM(nums []int) int {
	switch len(nums) {
	case 0:
		panic("can't do LCM on empty slice")
	case 1:
		return nums[0]
	case 2:
		return LCM2(nums[0], nums[1])
	default:
		return LCM2(nums[0], LCM(nums[1:]))
	}
}

// wrapper for the dumb types of math.Max
func MaxInt(a, b int) (max int) {
	return int(math.Max(float64(a), float64(b)))
}
