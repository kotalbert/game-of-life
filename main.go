package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Game struct {
	seed        int
	n           int
	generations int
	cells       [][]bool
}

func (g Game) ToString() any {

	var bld strings.Builder
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.cells[i][j] {
				bld.WriteRune('O')
			} else {
				bld.WriteRune(' ')
			}
		}
		bld.WriteRune('\n')
	}
	return bld.String()

}

// NextGeneration creates a new generation of cells
//
//	Create new game object with new cells
func (g Game) NextGeneration() *Game {

	newCells := createEmptyCells(g.n)

	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			neighbors := g.countNeighbors(i, j)
			switch neighbors {
			case 2: // cell lives if alive, or stays dead
				newCells[i][j] = g.cells[i][j]
			case 3: // cell is reborn if dead or lives on
				newCells[i][j] = true
			default: // cell dies
				newCells[i][j] = false
			}
		}

	}
	g.cells = newCells
	return &g
}

func maxIndex(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minIndex(a int, b int) int {
	if a < b {
		return a
	}
	return b

}

func (g Game) countNeighbors(i int, j int) int {

	// wrap around the edges
	iMinus := maxIndex(i-1, g.n-1)
	jMinus := maxIndex(j-1, g.n-1)
	iPlus := minIndex(j+1, g.n-1)
	jPlus := minIndex(i+1, g.n-1)

	neighbours := []bool{
		g.cells[iMinus][jMinus], g.cells[iMinus][j], g.cells[iMinus][iPlus],
		g.cells[i][jMinus], g.cells[i][iPlus],
		g.cells[jPlus][jMinus], g.cells[jPlus][j], g.cells[jPlus][iPlus]}
	var n int
	for _, v := range neighbours {
		if v {
			n++
		}
	}
	return n
}

func NewGame(seed int, n int, generations int) *Game {

	game := &Game{seed: seed, n: n, generations: generations}

	rand.Seed(int64(seed))

	game.cells = make([][]bool, n)

	initializeUniverse(n, game)

	return game
}

func createEmptyCells(n int) [][]bool {
	cells := make([][]bool, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]bool, n)
	}
	return cells
}

// initializeUniverse creates a 2D slice of bool,
//
//	the initial state of the universe
func initializeUniverse(n int, game *Game) {

	game.cells = createEmptyCells(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 50% chance of be, if rolls 1 then "alive"
			game.cells[i][j] = rand.Intn(2) == 1
		}
	}
}

func main() {

	n := readInt()
	seed := readInt()
	generations := readInt()
	game := NewGame(seed, n, generations)

	for i := 0; i < game.generations; i++ {
		fmt.Println(game.ToString())
		game = game.NextGeneration()

	}

}

func readInt() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
