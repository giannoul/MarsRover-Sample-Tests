package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiveSingleCommandShouldMove(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 9}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}
	CurrentGrid = NewGrid(10, 10)
	//grid.UpdateRoverPosition(marsRover.heading, marsRover.position)
	//grid.Draw()

	commands := []Command{B}
	marsRover.acceptCommands(commands)
	//grid.UpdateRoverPosition(marsRover.heading, marsRover.coordinates())
	//grid.Draw()
	expectedPosition := Coordinates{1, 8}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}
