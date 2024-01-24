package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for t := j + 1; t <= '9'; t++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(t)
				if i != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
