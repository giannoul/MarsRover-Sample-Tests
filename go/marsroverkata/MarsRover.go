package marsrover

type Coordinates struct {
	x int
	y int
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

type Command int

const (
	B Command = iota
	F
	L
	R
)

func (c Command) String() string {
	return [...]string{"B", "F", "L", "R"}[c]
}

type Obstacle struct {
	position Coordinates
}

type Plateau struct {
	maxX      int
	maxY      int
	obstacles []Obstacle
}

type Status int

const (
	OK Status = iota
	NOK
)

func (s Status) String() string {
	return [...]string{"OK", "NOK"}[s]
}

type MarsRover struct {
	plateau  Plateau
	heading  Direction
	position Coordinates
	status   Status
}

func (r *MarsRover) turnLeft() {
	// [0, 1, 2, 3]
	// [N, E, S, W]
	newDirection := W
	if r.heading-1 > 0 {
		newDirection = r.heading - 1
	}
	r.heading = newDirection
}

func (r MarsRover) currentLocation() interface{} {
	return ""
}

func (r MarsRover) acceptCommands(commands []Command) {

}

func (r MarsRover) coordinates() Coordinates {
	return Coordinates{0, 0}
}

func (r MarsRover) forward() {

}

func (r MarsRover) backward() {

}

func (r *MarsRover) turnRight() {
	// [0, 1, 2, 3]
	// [N, E, S, W]
	newDirection := N
	if r.heading < 3 {
		newDirection = r.heading + 1
	}
	r.heading = newDirection
}
