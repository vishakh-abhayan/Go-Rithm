package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}
	steps := 0
	for n != 1 {
		if n&1 == 1 {
			n = (n << 1) + n + 1
		} else {
			n >>= 1
		}
		steps++
	}
	return steps, nil
}
