package main


// Prints an array in spiral
func spiralPrint(spiral [][]int ) (out []int){
	// print top edge
	// do it in one loop
	if len(spiral) < 1 {
		return out
	}
	
	out = make([]int, len(spiral) * len(spiral[0]))

	

	for i, row := range spiral {
		for j, val := range row {

		}
	}

}

func main () {}