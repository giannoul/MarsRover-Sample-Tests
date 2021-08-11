package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoverWithObstacles(t *testing.T) {
	obstacles := []Obstacle{
		{position: Coordinates{2, 3}},
		{position: Coordinates{8, 8}},
		{position: Coordinates{9, 2}},
	}
	plateau := Plateau{maxX: 10, maxY: 10, obstacles: obstacles}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	CurrentGrid = NewGrid(10, 10)
	CurrentGrid.AddObstacles(plateau.obstacles)

	commands := []Command{B, F, L, F, F, R}
	marsRover.acceptCommands(commands)

	assert.Equal(t, "10 2 W", marsRover.currentLocation())
}

func TestRoverWithMoreObstacles(t *testing.T) {
	obstacles := []Obstacle{
		{position: Coordinates{2, 3}},
		{position: Coordinates{8, 8}},
		{position: Coordinates{9, 2}},
		{position: Coordinates{5, 5}},
		{position: Coordinates{6, 8}},
		{position: Coordinates{4, 2}},
	}
	plateau := Plateau{maxX: 10, maxY: 10, obstacles: obstacles}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	CurrentGrid = NewGrid(10, 10)
	CurrentGrid.AddObstacles(plateau.obstacles)

	commands := []Command{B, F, L, F, R, F, F, F, L, F, F, F, F, F, F, F, B}
	marsRover.acceptCommands(commands)

	assert.Equal(t, "6 5 W", marsRover.currentLocation())
}
