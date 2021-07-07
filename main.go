package main

import (
	"fmt"
	"math/rand"
	"time"
)

const gridsize = 9

var words = []string{"plan", "top", "football"}

func main() {

	var grid [gridsize][gridsize]string
	populateGrid(&grid)

	for i := 0; i < len(words); i++ {
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(len(grid[i]))
		col := rand.Intn(len(grid))
		insertWordHorizontally(words[i], row, &grid)
		insertWordVertically(words[i], col, &grid)
	}

	for i := 0; i < len(grid)-1; i++ {
		fmt.Println(grid[i])
	}

}

func insertWordVertically(word string, col int, grid *[9][9]string) {
	for n := 0; n < len(words)-1; n++ {
		for g := 0; g < len(grid[n]); g++ {
			if g > len(word)-1 {
				break
			}
			grid[g][col] = string(word[g])
		}

	}
}

func insertWordHorizontally(word string, row int, grid *[9][9]string) {
	for n := 0; n < len(words)-1; n++ {
		for g := 0; g < len(grid[n]); g++ {
			if g > len(word)-1 {
				break
			}
			grid[row][g] = string(word[g])
		}

	}
}

func populateGrid(grid *[gridsize][gridsize]string) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			randomchar := getRandomChar()
			grid[i][j] = randomchar
		}

	}
}

func getRandomChar() string {
	rand.Seed(time.Now().UnixNano())
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	randomchar := chars[rand.Intn(13)]
	return randomchar
}
