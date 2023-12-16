package main
import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)
func main() {
	if vaild() == true {
		array := [9][9]rune{}
		for i := 1; i <= 9; i++ {
			x := 0
			for _, c := range os.Args[i] {
				array[i-1][x] = c
				x++
			}
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if array[i][j] >= '1' && array[i][j] <= '9' {
					if Isdubl(array, array[i][j], i, j) == false {
						fmt.Println("Error")
						return
					}
				}
			}
		}
		if solve(array) == true {
			array = solve2(array)
			array = solve2(array)
			print(array)
		} else {
			fmt.Println("Error")
		}
	}
}
func vaild() bool {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return false
	} else {
		for i := 1; i <= 9; i++ {
			if len(os.Args[i]) != 9 {
				fmt.Println("Error")
				return false
			}
		}
		for i := 1; i <= 9; i++ {
			for _, z := range os.Args[i] {
				if z != '.' {
					if !(z >= '1' && z <= '9') {
						fmt.Println("Error")
						return false
					}
				}
			}
		}
	}
	return true
}
func print(array [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune((array[i][j]))
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		fmt.Print("\n")
	}
}
func inRow(array [9][9]rune, num rune, row int) bool {
	for i := 0; i < 9; i++ {
		if array[row][i] == num {
			return false
		}
	}
	return true
}
func inCol(array [9][9]rune, num rune, col int) bool {
	for i := 0; i < 9; i++ {
		if array[i][col] == num {
			return false
		}
	}
	return true
}
func inBox(array [9][9]rune, num rune, col int, row int) bool {
	row = row - row%3
	col = col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if array[row+i][col+j] == num {
				return false
			}
		}
	}
	return true
}
func isTrue(array [9][9]rune, num rune, row int, col int) bool {
	if inBox(array, num, col, row) && inCol(array, num, col) && inRow(array, num, row) {
		return true
	}
	return false
}
func solve(array [9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if array[i][j] == '.' {
				for v := '1'; v <= '9'; v++ {
					if isTrue(array, v, i, j) == true {
						array[i][j] = v
						if solve(array) == true {
							return true
						} else {
							array[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
func solve2(array [9][9]rune) [9][9]rune {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if array[i][j] == '.' {
				for v := '1'; v <= '9'; v++ {
					if isTrue(array, v, i, j) == true {
						array[i][j] = v
						if solve(array) != true {
							array[i][j] = '.'
						}
					}
				}
			}
		}
	}
	return array
}
func Isdubl(array [9][9]rune, num rune, row int, col int) bool {
	for i := 0; i < 9; i++ {
		if i != col {
			if array[row][i] == num {
				return false
			}
		}
	}
	for i := 0; i < 9; i++ {
		if i != row {
			if array[i][col] == num {
				return false
			}
		}
	}
	x := row - row%3
	z := col - col%3
	for i := x; i < 3; i++ {
		for j := z; j < 3; j++ {
			if row != i && col != j {
				if array[i][j] == num {
					return false
				}
			}
		}
	}
	return true
}
