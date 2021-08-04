package marsrover

import (
	"fmt"
	"strings"
)

type Grid struct {
	Height int
	Width  int
	Tiles  [][]string
}

func NewGrid(h int, w int) *Grid {
	g := Grid{Height: h, Width: w}
	g.initializeTiles()
	return &g
}

func (g Grid) DrawHorizontalLine() {
	fmt.Println(strings.Repeat("----", g.Width) + "---")
}

func (g Grid) DrawHorizontalNumbers() {
	fmt.Printf("  ")
	for i := 0; i < g.Width; i++ {
		fmt.Printf(" %2d ", i+1)
	}
	fmt.Println()
}

func (g *Grid) initializeTiles() {
	// initialize
	g.Tiles = make([][]string, g.Height)
	for i := 0; i < g.Height; i++ {
		g.Tiles[i] = make([]string, g.Width)
	}
	// fill with spaces
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			g.Tiles[i][j] = " "
		}
	}
}

func (g *Grid) Draw() {
	g.DrawHorizontalNumbers()
	for i := 0; i < g.Height; i++ {
		g.DrawHorizontalLine()
		for j := 0; j < g.Width; j++ {
			if j == 0 {
				fmt.Printf("%2d", i+1)
			}
			fmt.Printf("| %s ", g.Tiles[i][j])
		}
		fmt.Println("|")
	}
	g.DrawHorizontalLine()
}
