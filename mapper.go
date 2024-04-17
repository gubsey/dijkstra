package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strings"
)

type Grid [][]int

func (grid Grid) Print(steps []Crd) {
	for i := range len(grid) {
		var line strings.Builder
		for j := range len(grid[0]) {
			//sb := strings.Builder{}

			if i%2 == j%2 && i%2 == 1 {
				fmt.Fprintf(&line, "%-4s", "")
			} else if grid[i][j] == math.MaxInt {
				fmt.Fprintf(&line, "%-4s", "\u221e")
			} else {
				is_path := slices.Contains(steps, Crd{j, i})
				if is_path {
					fmt.Fprint(&line, "\u001b[32m")
				} else {
					fmt.Fprint(&line, "\u001b[31m")
				}
				fmt.Fprintf(&line, "%-4d", grid[i][j])

				fmt.Fprint(&line, "\u001b[0m")
			}
		}
		str := line.String()
		fmt.Println(str)
	}
}

/*
imagines a grid with extra cells inserted between each given cell.

a 3x3 will become a 5x5. a 2x1 becomes a 3x1
*/
func (Grid) gen(width int, height int, d_max int) (int, int, Grid) {
	grid := [][]int{}

	new_h := height + (height - 1)
	new_w := width + (width - 1)

	for h := range new_h {
		h_even := h%2 == 0

		row := []int{}
		for w := range new_w {
			w_even := w%2 == 0
			if h_even && w_even {
				row = append(row, 0)
			} else if h_even != w_even {
				row = append(row, (rand.Int()%d_max)+1)
			} else {
				row = append(row, 0)
			}
		}
		grid = append(grid, row)
	}

	return new_w, new_h, grid
}
