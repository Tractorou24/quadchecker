package quad

import (
	"os"

	"github.com/01-edu/z01"
)

func quadB() {
	args := os.Args[1:]

	if len(args) > 0 && len(args) < 3 {
		if IsNumeric(args[0]) && IsNumeric(args[1]) {
			QuadB(Atoi(args[0]), Atoi(args[1]))
		}
	}
}

func QuadB(x int, y int) {
	a := x
	x = y
	y = a

	//2D Array Definition
	var table = make([][]rune, x)
	for i := range table {
		table[i] = make([]rune, y)
	}

	//Complete Array

	for line := 0; line < x; line++ {
		for col := 0; col < y; col++ {
			if ((line == 0) || (line == x-1)) && ((col > 0) && (col < y-1)) {
				table[line][col] = 42
			} else if ((col == 0) || (col == y-1)) && ((line > 0) && (line < x-1)) {
				table[line][col] = 42
			} else if ((line == 0) && (col == 0)) || ((line == x-1) && (col == y-1)) {
				table[line][col] = 47
			} else if ((line == 0) && (col == y-1)) || ((line == x-1) && (col == 0)) {
				table[line][col] = 92
			} else {
				table[line][col] = 32
			}
		}
	}

	//Show Array
	for line := 0; line < x; line++ {
		for col := 0; col < y; col++ {
			z01.PrintRune(table[line][col])
		}
		z01.PrintRune('\n')
	}
}
