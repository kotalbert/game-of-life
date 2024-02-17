package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Game struct {
	seed  int
	n     int
	cells [][]bool
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

func NewGame(seed int, n int) *Game {

	g := &Game{seed: seed, n: n}

	rand.Seed(int64(seed))

	g.cells = make([][]bool, n)

	for i := 0; i < n; i++ {
		g.cells[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			// 50% chance of be, if rolls 1 then "alive"
			g.cells[i][j] = rand.Intn(2) == 1
		}
	}

	return g
}

func main() {

	n := readInt()
	seed := readInt()
	game := NewGame(seed, n)
	fmt.Println(game.ToString())

}

func readInt() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
