package main

import (
	"fmt"
)

type Position struct {
	x, y int
}

func (p Position) next(direction Direction) Position {
	switch direction {
	case North:
		return Position{p.x, p.y + 1}
	case South:
		return Position{p.x, p.y - 1}
	case East:
		return Position{p.x + 1, p.y}
	case West:
		return Position{p.x - 1, p.y}

	default:
		return p
	}
}

type Direction int

const (
	North = iota
	East
	South
	West
)

func (d Direction) turnLeft() Direction {
	newPosition := d - 1

	if newPosition < North {
		return West
	}
	return newPosition
}

func (d Direction) turnRight() Direction {
	newPosition := d + 1

	if newPosition > West {
		return North
	}
	return newPosition
}

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Lawn struct {
	topRightPosition Position
}

func (l Lawn) isValid(position Position) bool {
	return position.x >= 0 && position.y >= 0 &&
		position.x <= l.topRightPosition.x &&
		position.y <= l.topRightPosition.y
}

type Mower struct {
	position  Position
	direction Direction
}

type MowerAndCommands struct {
	mower *Mower
	commands []string
}

func (m Mower) String() string {
	return fmt.Sprintf("position: %v, direction: %s", m.position, m.direction)
}

func (m Mower) positionAsString() string {
	return fmt.Sprintf("%d %d %s", m.position.x, m.position.y, m.direction)
}

func (m Mower) move(commands []string, lawn Lawn) *Mower {
	for _, command := range commands {
		switch command {

		case "A":
			newPosition := m.position.next(m.direction)
			if lawn.isValid(newPosition) {
				m.position = newPosition
			}
		case "G":
			m.direction = m.direction.turnLeft()

		case "D":
			m.direction = m.direction.turnRight()

		}
	}

	return &m
}
