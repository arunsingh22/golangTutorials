package main

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	var ans float64 = 1.0

	exp := int(math.Abs(float64(n)))
	for exp > 0 {
		// check if it's odd
		if (exp & 1) == 1 {
			ans *= x
			exp--
		} else {
			// exp is even
			exp /= 2 // exp >> 1
			x *= x
		}
	}

	if n < 0 {
		ans = 1.0 / ans
	}
	return ans
}
