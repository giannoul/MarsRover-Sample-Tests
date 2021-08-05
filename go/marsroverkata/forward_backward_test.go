package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleCommandForwardOneTileFacingNorth(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 9}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	marsRover.forward()

	expectedPosition := Coordinates{1, 8}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandForwardOneTileFacingEast(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 9}
	marsRover := MarsRover{plateau: plateau, heading: E, position: startingPosition}

	marsRover.forward()

	expectedPosition := Coordinates{2, 9}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandForwardOneTileFacingSouth(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 9}
	marsRover := MarsRover{plateau: plateau, heading: S, position: startingPosition}

	marsRover.forward()

	expectedPosition := Coordinates{1, 10}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandForwardOneTileFacingWest(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{2, 9}
	marsRover := MarsRover{plateau: plateau, heading: W, position: startingPosition}

	marsRover.forward()

	expectedPosition := Coordinates{1, 9}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandBackwardOneTileFacingNorth(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{5, 5}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	marsRover.backward()

	expectedPosition := Coordinates{5, 6}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandBackwardOneTileFacingEast(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{5, 5}
	marsRover := MarsRover{plateau: plateau, heading: E, position: startingPosition}

	marsRover.backward()

	expectedPosition := Coordinates{4, 5}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandBackwardOneTileFacingSouth(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{5, 5}
	marsRover := MarsRover{plateau: plateau, heading: S, position: startingPosition}

	marsRover.backward()

	expectedPosition := Coordinates{5, 4}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}

func TestSingleCommandBackwardOneTileFacingWest(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{5, 5}
	marsRover := MarsRover{plateau: plateau, heading: W, position: startingPosition}

	marsRover.backward()

	expectedPosition := Coordinates{6, 5}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}
