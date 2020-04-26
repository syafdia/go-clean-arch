package slice

func FilterInt(xs []int, f func(int) bool) []int {
	vs := []int{}

	for _, x := range xs {
		if f(x) {
			vs = append(vs, x)
		}
	}

	return vs
}
