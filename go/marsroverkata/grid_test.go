package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridVisualization(t *testing.T) {
	grid := NewGrid(10, 10)
	grid.Draw()
	assert.Equal(t, true, true, "We want to just see the output")
}
