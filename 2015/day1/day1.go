package day1

func Part1(s string) int {
	var res int
	for _, rn := range s {
		switch rn {
		case '(':
			res++
		case ')':
			res--
		}
	}
	return res
}

// Part2 returns -1 if the basement is never reached.
func Part2(s string) int {
	var res int

	for i, rn := range s {
		switch rn {
		case '(':
			res++
		case ')':
			res--
		}
		if res == -1 {
			return i + 1
		}
	}

	return -1
}
