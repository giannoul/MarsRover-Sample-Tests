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
	return r.position
}

func (r *MarsRover) forward() {
	curr := r.position
	switch r.heading {
	case N:
		r.position = Coordinates{curr.x, curr.y - 1}
	case E:
		r.position = Coordinates{curr.x + 1, curr.y}
	case S:
		r.position = Coordinates{curr.x, curr.y + 1}
	case W:
		r.position = Coordinates{curr.x - 1, curr.y}
	default:
		panic("I won't know which way I am heading!!")
	}
}

func (r *MarsRover) backward() {
	curr := r.position
	switch r.heading {
	case N:
		r.position = Coordinates{curr.x, curr.y + 1}
	case E:
		r.position = Coordinates{curr.x - 1, curr.y}
	case S:
		r.position = Coordinates{curr.x, curr.y - 1}
	case W:
		r.position = Coordinates{curr.x + 1, curr.y}
	default:
		panic("I won't know which way I am heading!!")
	}
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
