package marsrover

import "fmt"

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
	return fmt.Sprintf("%d %d %s", r.position.x, r.position.y, r.heading)
}

func (r *MarsRover) acceptCommands(commands []Command) {
	r.showOnGrid() // Show the initial position
	for i := range commands {
		if r.status == NOK { //The rover will stay still is an obstacle is found
			fmt.Println(" 💥 🚨 Obstacle ahead, aborting!!!! 🚀 ")
			break
		}
		switch commands[i] {
		case B:
			r.backward()
		case F:
			r.forward()
		case L:
			r.turnLeft()
		case R:
			r.turnRight()
		default:
			panic("I don't know that command!!")
		}
		r.showOnGrid()
	}

}

func (r *MarsRover) checkForObstacleCrash(c Coordinates) {
	for i := range r.plateau.obstacles {
		if c == r.plateau.obstacles[i].position {
			r.status = NOK
		}
	}
}

func (r MarsRover) coordinates() Coordinates {
	return r.position
}

func (r *MarsRover) fold() {
	if r.position.x > MaxTerrainPos {
		r.position.x = 1
	}
	if r.position.x < 1 {
		r.position.x = MaxTerrainPos
	}
	if r.position.y > MaxTerrainPos {
		r.position.y = 1
	}
	if r.position.y < 1 {
		r.position.y = MaxTerrainPos
	}
}

func (r *MarsRover) forward() {
	curr, newPosition := r.position, r.position

	switch r.heading {
	case N:
		newPosition = Coordinates{curr.x, curr.y + 1}
	case E:
		newPosition = Coordinates{curr.x + 1, curr.y}
	case S:
		newPosition = Coordinates{curr.x, curr.y - 1}
	case W:
		newPosition = Coordinates{curr.x - 1, curr.y}
	default:
		panic("I won't know which way I am heading!!")
	}
	r.checkForObstacleCrash(newPosition)
	if r.status == OK {
		r.position = newPosition
	}
	r.fold()
}

func (r *MarsRover) backward() {
	curr, newPosition := r.position, r.position
	switch r.heading {
	case N:
		newPosition = Coordinates{curr.x, curr.y - 1}
	case E:
		newPosition = Coordinates{curr.x - 1, curr.y}
	case S:
		newPosition = Coordinates{curr.x, curr.y + 1}
	case W:
		newPosition = Coordinates{curr.x + 1, curr.y}
	default:
		panic("I won't know which way I am heading!!")
	}

	r.checkForObstacleCrash(newPosition)
	if r.status == OK {
		r.position = newPosition
	}
	r.fold()
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

func (r *MarsRover) showOnGrid() {
	CurrentGrid.UpdateRoverPosition(r.heading, r.position)
	CurrentGrid.Draw()
	fmt.Println()
}
