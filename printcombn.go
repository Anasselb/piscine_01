package piscine

import "github.com/01-edu/z01"

func onecombs() {

	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
		if i != '9' {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')

}

func twocombs() {

	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i != '8' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

	z01.PrintRune('\n')

}

func threecombs() {
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

func fourcombs() {
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

						if i != '9' {
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

func fivecombs() {

}

func sixcombs() {
	for i := '0'; i < '5'; i++ {
		for j := i + 1; j < '6'; j++ {
			for k := j + 1; k < '7'; k++ {
				for l := k + 1; l < '8'; l++ {
					for m := l + 1; m < '9'; m++ {
						for n := m + 1; n <= '9'; n++ {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(m)
							z01.PrintRune(n)
							if i != '4' {
								z01.PrintRune(',')
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func sevencombs() {

	for i := '0'; i < '4'; i++ {
		for j := i + 1; j < '5'; j++ {
			for k := j + 1; k < '6'; k++ {
				for l := k + 1; l < '7'; l++ {
					for m := l + 1; m < '8'; m++ {
						for n := m + 1; n < '9'; n++ {
							for o := n + 1; n <= '9'; o++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(m)
								z01.PrintRune(n)
								z01.PrintRune(o)
								if i != '3' {
									z01.PrintRune(',')
								}
							}

						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func eightcombs() {

	for i := '0'; i < '3'; i++ {
		for j := i + 1; j < '4'; j++ {
			for k := j + 1; k < '5'; k++ {
				for l := k + 1; l < '6'; l++ {
					for m := l + 1; m < '7'; m++ {
						for n := m + 1; n < '8'; n++ {
							for o := n + 1; n < '9'; o++ {
								for p := o + 1; p <= '9'; p++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune(n)
									z01.PrintRune(o)
									z01.PrintRune(p)
									if i != '3' {
										z01.PrintRune(',')
									}
								}

							}

						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')

}

func ninecombs() {
	for i := '0'; i < '2'; i++ {
		for j := i + 1; j < '3'; j++ {
			for k := j + 1; k < '4'; k++ {
				for l := k + 1; l < '5'; l++ {
					for m := l + 1; m < '6'; m++ {
						for n := m + 1; n < '7'; n++ {
							for o := n + 1; n < '8'; o++ {
								for p := o + 1; p < '9'; p++ {
									for q := p + 1; q <= '9'; q++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune(o)
										z01.PrintRune(p)
										z01.PrintRune(q)
										if i != '2' {
											z01.PrintRune(',')
										}
									}

								}

							}

						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN(n int) {

	if n < 0 || n > 9 {
		return
	}

	if n == 1 {
		onecombs()
	} else if n == 2 {
		twocombs()
	} else if n == 3 {
		threecombs()
	} else if n == 4 {
		fourcombs()
	} else if n == 5 {
		fivecombs()
	} else if n == 6 {
		sixcombs()
	} else if n == 7 {
		sevencombs()
	} else if n == 8 {
		eightcombs()
	} else {
		ninecombs()
	}

}
