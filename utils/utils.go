package utils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
