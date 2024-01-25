package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	result := ""
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for n > 0 {
		digit := n % 10
		result = result + string(rune(digit+'0'))
		n = n / 10
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(rune(result[i]))
	}
}
