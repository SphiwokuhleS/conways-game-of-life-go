package domain

import (
	"fmt"
	"strings"
)

func GenerateNextGeneration(grid [21][21]int) [21][21]int {
	//grid size will be constant for now
	newGrid := [21][21]int{}

	// fmt.Println(grid[2+1][2+-1])
	//to avoid out of bounds the grid starts at 1 and ends at last element -1
	for row := 1; row < len(grid)-1; row++ {
		for column := 1; column < len(grid)-1; column++ {

			aliveNeighbours := 0
			for nRow := -1; nRow <= 1; nRow++ {
				for nCol := -1; nCol <= 1; nCol++ {
					aliveNeighbours += grid[row+nRow][column+nCol]
				}
			}

			//current position is added, so we remove its increment from aliveNeighbours
			aliveNeighbours -= grid[row][column]

			if grid[row][column] == 1 && aliveNeighbours <= 1 {
				newGrid[row][column] = 0

			} else if (grid[row][column] == 1) && aliveNeighbours >= 4 {
				newGrid[row][column] = 0

			} else if grid[row][column] == 0 && aliveNeighbours == 3 {
				newGrid[row][column] = 1

			} else {
				newGrid[row][column] = grid[row][column]
			}
		}
	}
	return newGrid
}

func MultipleGenerations(numberOfGenerations int, grid [21][21]int) [21][21]int {

	var newGrid = [21][21]int{}
	for i := 0; i <= numberOfGenerations; i++ {
		newGrid = GenerateNextGeneration(grid)
	}
	return newGrid
}

func ToString(grid [21][21]int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(grid)), " "), "")
}

func ToJsonArray(arrayString string) string {
	stripedGrid := strings.Replace(arrayString, " ", ",", -1)
	return stripedGrid
}
