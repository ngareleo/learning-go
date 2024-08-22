package main

import "fmt"

// Prints an array in spiral
func spiralPrint(spiral [][]int) (out []int) {
	// print top edge
	// do it in one loop

	if x := len(spiral); x < 1 {
		return out
	} else if y := len(spiral[0]); y < 1 {
		return out
	} else if x != y {
		return out
	}

	sizeY := len(spiral)
	sizeX := len(spiral[0])

	out = make([]int, sizeX*sizeY)

	// plan is to map (x, y) to an index to the final array
	for x, row := range spiral {
		for y, val := range row {
			pos := mapXYToCoordinate(x, y, sizeX, sizeY)
			out[pos] = val
		}
	}

	return
}

func mapXYToCoordinate(x, y, sizeX, sizeY int) int {
	return 0
}

func main() {
	arr := [][]int{
		{0, 1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10, 11},
		{12, 13, 14, 15, 16, 17},
		{18, 19, 20, 21, 22, 23},
	}
	spiral := spiralPrint(arr)
	fmt.Println("Printing", arr, "in a spiral")
	fmt.Println(spiral)
}
