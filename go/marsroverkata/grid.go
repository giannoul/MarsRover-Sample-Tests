package marsrover

import (
	"fmt"
	"strings"
)

var Markers = [4]string{"👆", "👉", "👇", "👈"}

var MaxTerrainPos = 10
var CurrentGrid *Grid

type Grid struct {
	Height int
	Width  int
	Tiles  [][]string
}

func NewGrid(h int, w int) *Grid {
	MaxTerrainPos = h
	g := Grid{Height: h, Width: w}
	g.initializeTiles()
	return &g
}

func (g Grid) AddObstacles(obstacles []Obstacle) {
	for i := range obstacles {
		p := obstacles[i].position
		// The -1 is because on slices we start counting from 0
		g.Tiles[p.y-1][p.x-1] = "🚧"
	}

}

func (g Grid) UpdateRoverPosition(h Direction, p Coordinates) {
	// Clear the previous rover position
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if g.Tiles[i][j] == "👆" || g.Tiles[i][j] == "👉" || g.Tiles[i][j] == "👇" || g.Tiles[i][j] == "👈" {
				g.Tiles[i][j] = "📍"
			}
		}
	}

	g.Tiles[p.y-1][p.x-1] = Markers[h]
}

func (g Grid) DrawHorizontalLine() {
	fmt.Println("  " + strings.Repeat("-----", g.Width) + "---")
}

func (g Grid) DrawHorizontalNumbers() {
	fmt.Printf("    ")
	for i := 0; i < g.Width; i++ {
		fmt.Printf("  %2d ", i+1)
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
			g.Tiles[i][j] = "  "
		}
	}
}

func (g *Grid) Draw() {
	// draw North
	fmt.Println(strings.Repeat(" ", g.Width*2) + "  N")

	for i := g.Width - 1; i >= 0; i-- {
		g.DrawHorizontalLine()
		for j := 0; j < g.Height; j++ {
			// draw West
			if i == (g.Height/2)-1 && j == 0 {
				fmt.Printf("W %2d", i+1)
			} else {
				if j == 0 {
					fmt.Printf("  %2d", i+1)
				}
			}

			fmt.Printf("| %s ", g.Tiles[i][j])
		}
		// draw East
		if i == (g.Height/2)-1 {
			fmt.Println("|  E")
		} else {
			fmt.Println("|")
		}
	}
	g.DrawHorizontalLine()
	g.DrawHorizontalNumbers()
	// draw South
	fmt.Println(strings.Repeat(" ", g.Width*2) + "  S")
}
