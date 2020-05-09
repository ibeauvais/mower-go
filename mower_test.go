package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirection_turnLeft(t *testing.T) {
	tests := []struct {
		name           string
		inputDirection Direction
		want           Direction
	}{
		{
			name:           "From North direction",
			inputDirection: North,
			want:           West,
		},
		{
			name:           "From West direction",
			inputDirection: West,
			want:           South,
		},
		{
			name:           "From South direction",
			inputDirection: South,
			want:           East,
		},
		{
			name:           "From East direction",
			inputDirection: East,
			want:           North,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.inputDirection.turnLeft()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestDirection_turnRight(t *testing.T) {
	tests := []struct {
		name           string
		inputDirection Direction
		want           Direction
	}{
		{
			name:           "From North direction",
			inputDirection: North,
			want:           East,
		},
		{
			name:           "From West direction",
			inputDirection: West,
			want:           North,
		},
		{
			name:           "From South direction",
			inputDirection: South,
			want:           West,
		},
		{
			name:           "From East direction",
			inputDirection: East,
			want:           South,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.inputDirection.turnRight()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name           string
		inputDirection Direction
		want           string
	}{
		{
			name:           "From North direction",
			inputDirection: North,
			want:           "North",
		},
		{
			name:           "From West direction",
			inputDirection: West,
			want:           "West",
		},
		{
			name:           "From South direction",
			inputDirection: South,
			want:           "South",
		},
		{
			name:           "From East direction",
			inputDirection: East,
			want:           "East",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.inputDirection.String()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestLawn_isValid(t *testing.T) {

	tests := []struct {
		name             string
		topRightPosition Position
		positionToValid  Position
		want             bool
	}{
		{
			name:             "valid position center",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{2, 3},
			want:             true,
		},
		{
			name:             "valid position top right",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{5, 5},
			want:             true,
		},
		{
			name:             "valid position bottom left",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{0, 0},
			want:             true,
		},
		{
			name:             "Invalid position x outside left",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{-1, 0},
			want:             false,
		},
		{
			name:             "Invalid position y outside bottom",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{0, -1},
			want:             false,
		},
		{
			name:             "Invalid position x outside right",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{6, 0},
			want:             false,
		},
		{
			name:             "Invalid position y outside top",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{0, 6},
			want:             false,
		},
		{
			name:             "Invalid position x and y",
			topRightPosition: Position{5, 5},
			positionToValid:  Position{6, 6},
			want:             false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := Lawn{
				topRightPosition: test.topRightPosition,
			}
			got := l.isValid(test.positionToValid)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPosition_next(t *testing.T) {

	tests := []struct {
		name      string
		direction Direction
		position  Position
		want      Position
	}{
		{
			name:      "From North direction",
			direction: North,
			position:  Position{1, 1},
			want:      Position{1, 2},
		},
		{
			name:      "From West direction",
			direction: West,
			position:  Position{1, 1},
			want:      Position{0, 1},
		},
		{
			name:      "From South direction",
			direction: South,
			position:  Position{1, 1},
			want:      Position{1, 0},
		},
		{
			name:      "From East direction",
			direction: East,
			position:  Position{1, 1},
			want:      Position{2, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.position.next(test.direction)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMower_move(t *testing.T) {
	tests := []struct {
		name     string
		mower    Mower
		commands []string
		lawn     Lawn
		want     Mower
	}{
		{
			name:     "First case",
			mower:    Mower{position: Position{1, 2}, direction: North},
			commands: []string{"G", "A", "G", "A", "G", "A", "G", "A", "A"},
			lawn:     Lawn{topRightPosition: Position{5, 5}},
			want:     Mower{position: Position{1, 3}, direction: North},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mower := test.mower.move(test.commands, test.lawn)
			assert.Equal(t, test.want, *mower)
		})
	}
}
