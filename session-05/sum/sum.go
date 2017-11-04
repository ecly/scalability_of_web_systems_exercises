package sum

// All returns the sum of the given values.
func All(vs ...int) int {
	return recursive(vs)
}

func recursive(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return vs[0] + recursive(vs[1:])
}
