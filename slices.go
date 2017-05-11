package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		row := make([]uint8, dx)
		for j := range row {	
			row[j] = uint8((i + j) / 2)
		}
		result[i] = row
	}
	return result
}

func main() {
	pic.Show(Pic)
}
