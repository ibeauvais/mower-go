package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseTopRightPosition(t *testing.T) {

	tests := []struct {
		name    string
		line    string
		want    *Lawn
		wantErr bool
	}{
		{
			name:    "true format '5 5'",
			line:    "5 5",
			want:    &Lawn{Position{5, 5}},
			wantErr: false,
		},
		{
			name:    "bad format '55'",
			line:    "55",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "bad format '5 T'",
			line:    "5 T",
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseTopRightPosition(test.line)

			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want, got)

		})
	}
}

func Test_parseMowers(t *testing.T) {
	tests := []struct {
		name    string
		lines   []string
		want    *[]MowerAndCommands
		wantErr bool
	}{
		{
			name:  "one mower",
			lines: []string{"1 2 N", "GAGAGAGAA"},
			want: &[]MowerAndCommands{
				{
					&Mower{
						Position{1, 2},
						North,
					},
					[]string{"G", "A", "G", "A", "G", "A", "G", "A", "A"},
				},
			},
			wantErr: false,
		},
		{
			name:  "two mower",
			lines: []string{"1 2 N", "GAGAGAGAA", "3 3 E", "AADAADADDA"},
			want: &[]MowerAndCommands{
				{
					&Mower{
						Position{1, 2},
						North,
					},
					[]string{"G", "A", "G", "A", "G", "A", "G", "A", "A"},
				},
				{
					&Mower{
						Position{3, 3},
						East,
					},
					[]string{"A", "A", "D", "A", "A", "D", "A", "D", "D", "A"},
				},
			},
			wantErr: false,
		},
		{
			name:    "Error in first line",
			lines:   []string{"1 2N", "GAGAGAGAA"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseMowers(test.lines)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want, got)
		})
	}
}
