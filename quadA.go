package quad

import (
	"os"

	"github.com/01-edu/z01"
)

func quadA() {
	args := os.Args[1:]

	if len(args) > 0 && len(args) < 3 {
		if IsNumeric(args[0]) && IsNumeric(args[1]) {
			QuadA(Atoi(args[0]), Atoi(args[1]))
		}
	}
}

func QuadA(x int, y int) {
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
				table[line][col] = 45
			} else if ((col == 0) || (col == y-1)) && ((line > 0) && (line < x-1)) {
				table[line][col] = 124
			} else if ((line == 0) && (col == 0)) || ((line == 0) && (col == y-1)) || ((line == x-1) && (col == 0)) || ((line == x-1) && (col == y-1)) {
				table[line][col] = 111
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

func Atoi(s string) int {
	strTable := []rune(s)
	var table = make([]int, len(strTable))
	var res int
	var isNegative bool

	for i := 0; i < len(strTable); i++ {
		table[i] = int(strTable[i] - 48)
	}
	for j := 0; j < len(strTable); j++ {
		if (table[j] == 0) || (table[j] == 1) || (table[j] == 2) || (table[j] == 3) || (table[j] == 4) || (table[j] == 5) || (table[j] == 6) || (table[j] == 7) || (table[j] == 8) || (table[j] == 9) || (table[0] == 43-48) || (table[0] == 45-48) {
			if table[0] == 45-48 {
				isNegative = true
				table[0] = 0
			}
			if table[0] == 43-48 {
				table[0] = 0
			}
			if len(s) > 2 {
				if (table[1] == 43-48) || (table[1] == 45-48) {
					return 0
				}
			}
			res = res*10 + table[j]
		} else {
			return 0
		}
	}
	if isNegative == true {
		return res * -1
	} else {
		return res
	}
}

func IsNumeric(s string) bool {
	tab := []rune(s)
	var isValid bool

	for i := 0; i < len(tab); i++ {
		if tab[i] > 47 && tab[i] < 58 {
			isValid = true
		} else {
			isValid = false
		}
		if isValid == false {
			return false
		}
	}
	return true
}
