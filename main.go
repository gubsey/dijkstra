package main

import (
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	argw, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("invalid width")
	}
	argh, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("invalid height")
	}

	width, height, grid := Grid{}.gen(argw, argh, 10)

	for i := range height {
		for j := range width {
			if i%2 == 0 && j%2 == 0 {
				grid[i][j] = math.MaxInt
			}
		}
	}
	grid[0][0] = 0

	var ol OList
	ol.Insert(0, 0, 0, nil)
	var end *Node

	for {
		node, err := ol.Pop()
		if err != nil {
			log.Fatalf("%e", err)
		}

		if node.x+1 == width && node.y+1 == height {
			grid[node.y][node.x] = node.val
			//node.steps = append(node.steps, Crd{node.x, node.y})
			end = node
			break
		}

		//skip nodes that are greater that their cooresponding cell
		if grid[node.y][node.x] < node.val {
			continue
		} else {
			grid[node.y][node.x] = node.val
		}

		// right
		DirOp(node.crd(), Crd{node.x + 2, node.y}, Crd{node.x + 1, node.y}, node, grid, &ol)

		// left
		DirOp(node.crd(), Crd{node.x - 2, node.y}, Crd{node.x - 1, node.y}, node, grid, &ol)

		// down
		DirOp(node.crd(), Crd{node.x, node.y + 2}, Crd{node.x, node.y + 1}, node, grid, &ol)

		// up
		DirOp(node.crd(), Crd{node.x, node.y - 2}, Crd{node.x, node.y - 1}, node, grid, &ol)
	}

	steps := []Crd{end.crd()}
	for end.prev != nil {
		x_mid := (end.x + end.prev.x) / 2
		y_mid := (end.y + end.prev.y) / 2
		mid := Crd{x_mid, y_mid}
		end = end.prev
		steps = append(steps, mid)
		steps = append(steps, end.crd())
	}
	grid.Print(steps)
}

func DirOp(src, dst, mid Crd, node *Node, grid [][]int, ol *OList) {
	if dst.x < 0 || dst.x >= len(grid[0]) || dst.y < 0 || dst.y >= len(grid) {
		return

	}
	dest_val := grid[dst.y][dst.x]
	new_val := node.val + grid[mid.y][mid.x]
	if new_val < dest_val {
		ol.Insert(dst.x, dst.y, new_val, node)
	}
}
