package basics

func fibonaci() func() int {
	prev := [2]int{0, 1}
	iter := 0

	return func() int {
		if iter < 2 {
			next := prev[iter]
			iter += 1

			return next
		}

		next := prev[0] + prev[1]
		prev[0] = prev[1]
		prev[1] = next

		return next
	}
}

func fibo_recursive(x int) int {
	if x < 2 {
		return x
	}

	return fibo_recursive(x-1) + fibo_recursive(x-2)
}
