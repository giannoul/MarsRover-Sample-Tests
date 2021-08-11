package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridWithRoverVisualization(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	grid := NewGrid(10, 10)
	grid.UpdateRoverPosition(marsRover.heading, marsRover.position)
	grid.Draw()

	marsRover.turnRight()
	grid.UpdateRoverPosition(marsRover.heading, marsRover.position)
	grid.Draw()
	assert.Equal(t, true, true, "We want to just see the output")
}
