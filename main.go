package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Game struct {
	n     int
	cells [][]bool
}

func (g Game) ToString() string {

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

func (g Game) getNumberOfAliveCells() int {
	var alive int
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.cells[i][j] {
				alive++
			}
		}
	}
	return alive
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

func (g Game) countNeighbors(i int, j int) int {

	// wrap around the edges, flip to the other side if out of bounds
	nw := g.cells[(i-1+g.n)%g.n][(j-1+g.n)%g.n]
	no := g.cells[(i-1+g.n)%g.n][j]
	ne := g.cells[(i-1+g.n)%g.n][(j+1)%g.n]
	we := g.cells[i][(j-1+g.n)%g.n]
	ea := g.cells[i][(j+1)%g.n]
	sw := g.cells[(i+1)%g.n][(j-1+g.n)%g.n]
	so := g.cells[(i+1)%g.n][j]
	se := g.cells[(i+1)%g.n][(j+1)%g.n]

	neighbours := []bool{
		nw, no, ne,
		we, ea,
		sw, so, se,
	}
	var n int
	for _, v := range neighbours {
		if v {
			n++
		}
	}
	return n
}

func NewGame(n int) *Game {

	game := &Game{n: n}

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
	game := NewGame(n)

	for i := 1; i <= 10; i++ {
		game = game.NextGeneration()
		fmt.Printf("Generation #%d\nAlive: %d\n", i, game.getNumberOfAliveCells())
		fmt.Println(game.ToString())
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
