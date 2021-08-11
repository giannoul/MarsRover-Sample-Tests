package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowInstructions(t *testing.T) {
	plateau := Plateau{maxX: 10, maxY: 10}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}
	CurrentGrid = NewGrid(10, 10)

	commands := []Command{B, F, L, F, F, R}
	marsRover.acceptCommands(commands)

	assert.Equal(t, "9 2 N", marsRover.currentLocation())
}
