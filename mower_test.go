package main

import (
	"reflect"
	"testing"
)

func TestDirection_turnLeft(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want Direction
	}{
		{
			name: "From North direction",
			d:    North,
			want: West,
		},
		{
			name: "From West direction",
			d:    West,
			want: South,
		},
		{
			name: "From South direction",
			d:    South,
			want: East,
		},
		{
			name: "From East direction",
			d:    East,
			want: North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.turnLeft(); got != tt.want {
				t.Errorf("turnLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_turnRight(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want Direction
	}{
		{
			name: "From North direction",
			d:    North,
			want: East,
		},
		{
			name: "From West direction",
			d:    West,
			want: North,
		},
		{
			name: "From South direction",
			d:    South,
			want: West,
		},
		{
			name: "From East direction",
			d:    East,
			want: South,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.turnRight(); got != tt.want {
				t.Errorf("turnRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want string
	}{
		{
			name: "From North direction",
			d:    North,
			want: "North",
		},
		{
			name: "From West direction",
			d:    West,
			want: "West",
		},
		{
			name: "From South direction",
			d:    South,
			want: "South",
		},
		{
			name: "From East direction",
			d:    East,
			want: "East",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lawn{
				topRightPosition: tt.topRightPosition,
			}
			if got := l.isValid(tt.positionToValid); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.position
			if got := p.next(tt.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
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
			lawn:     Lawn{topRightPosition:Position{5,5}},
			want:     Mower{position: Position{1, 3}, direction: North},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mower
			mower := m.move(tt.commands, tt.lawn)
			if !reflect.DeepEqual(*mower, tt.want) {
				t.Errorf("next() = %v, want %v", m, tt.want)
			}
		})
	}
}
