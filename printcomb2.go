package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for t := '0'; t <= '9'; t++ {
				for v := '0'; v <= '9'; v++ {
					if (i < t) || (i == t && j < v) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(t)
						z01.PrintRune(v)

						if i != '9' || j != '8' || t != '9' || v != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
