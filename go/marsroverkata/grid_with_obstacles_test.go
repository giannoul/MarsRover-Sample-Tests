package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridWithObstaclesVisualization(t *testing.T) {
	obstacles := []Obstacle{
		Obstacle{position: Coordinates{2, 3}},
		Obstacle{position: Coordinates{8, 8}},
	}
	plateau := Plateau{maxX: 5, maxY: 5, obstacles: obstacles}
	grid := NewGrid(10, 10)
	grid.AddObstacles(plateau.obstacles)
	grid.Draw()
	assert.Equal(t, true, true, "We want to just see the output")
}
