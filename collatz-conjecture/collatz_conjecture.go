package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}

	for i := 0; ; i++ {
		if n == 1 {
			return i, nil
		}

		n = CollatzStep(n)
	}
}

func CollatzStep(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
