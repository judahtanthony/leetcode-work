package main

import "fmt"

func setZeroes(matrix [][]int) {
	rown := len(matrix)
	if rown == 0 {
		return
	}
	coln := len(matrix[0])
	firstRow, firstCol := false, false
	for row := 0; row < rown; row++ {
		for col := 0; col < coln; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
				if row == 0 {
					firstRow = true
				}
				if col == 0 {
					firstCol = true
				}
			}
		}
	}
	for row := 1; row < rown; row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < coln; col++ {
				matrix[row][col] = 0
			}
		}
	}
	for col := 1; col < coln; col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < rown; row++ {
				matrix[row][col] = 0
			}
		}
	}
	if firstRow {
		for col := 0; col < coln; col++ {
			matrix[0][col] = 0
		}
	}
	if firstCol {
		for row := 0; row < rown; row++ {
			matrix[row][0] = 0
		}
	}
}

func main() {
	m := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	setZeroes(m)
	fmt.Println(m)
}
