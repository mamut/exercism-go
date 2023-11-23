package luhn

import (
	"strconv"
)

func Valid(id string) bool {
	digits, err := toDigits(id)

	if len(digits) <= 1 {
		return false
	}

	if err != nil {
		return false
	}

	for i := 0; i < len(digits); i++ {
		if i%2 == 0 {
			continue
		}

		offset := len(digits) - i - 1

		newDigit := digits[offset] * 2
		if newDigit > 9 {
			newDigit -= 9
		}

		digits[offset] = newDigit
	}

	sum := 0
	for _, digit := range digits {
		sum += digit
	}

	return sum%10 == 0
}

func toDigits(id string) ([]int, error) {
	digits := []int{}

	for _, ch := range id {
		if ch == ' ' {
			continue
		}

		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			return nil, err
		}

		digits = append(digits, digit)
	}

	return digits, nil
}
