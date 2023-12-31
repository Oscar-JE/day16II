package matrix

import "fmt"

type Matrix struct {
	m [][]int
}

func Init(nrRows int, nrCols int) Matrix {
	m := [][]int{}
	for i := 0; i < nrRows; i++ {
		row := []int{}
		for j := 0; j < nrCols; j++ {
			row = append(row, 0)
		}
		m = append(m, row)
	}
	return Matrix{m: m}
}

func (m Matrix) String() string {
	rep := ""
	for i := range m.m {
		rep += fmt.Sprintln(m.m[i])
	}
	return rep
}

func (m *Matrix) Set(row int, col int, value int) {
	m.m[row][col] = value
}

func (m Matrix) Get(row int, col int) int {
	return m.m[row][col]
}
