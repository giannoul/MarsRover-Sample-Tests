package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanRotateRight1(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	marsRover.turnRight()

	assert.Equal(t, E, marsRover.heading)
}

func TestCanRotateRight2(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: W, position: startingPosition}

	marsRover.turnRight()

	assert.Equal(t, N, marsRover.heading)
}
