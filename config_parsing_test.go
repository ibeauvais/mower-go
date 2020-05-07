package main

import (
	"reflect"
	"testing"
)

func Test_parseTopRightPosition(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    *Lawn
		wantErr bool
	}{
		{
			name:    "true format '5 5'",
			args:    struct{ line string }{line: "5 5"},
			want:    &Lawn{Position{5, 5}},
			wantErr: false,
		},
		{
			name:    "bad format '55'",
			args:    struct{ line string }{line: "55"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "bad format '5 T'",
			args:    struct{ line string }{line: "5 T"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseTopRightPosition(test.args.line)
			if (err != nil) != test.wantErr {
				t.Errorf("parseTopRightPosition() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("parseTopRightPosition() got = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_parseMowers(t *testing.T) {
	type args struct {
		lines []string
	}

	tests := []struct {
		name    string
		args    args
		want    *[]MowerAndCommands
		wantErr bool
	}{
		{
			name: "one mower",
			args: struct{ lines []string }{lines: []string{"1 2 N", "GAGAGAGAA"}},
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
			name: "two mower",
			args: struct{ lines []string }{lines: []string{"1 2 N", "GAGAGAGAA", "3 3 E", "AADAADADDA"}},
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
			args:    struct{ lines []string }{lines: []string{"1 2N", "GAGAGAGAA"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMowers(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMowers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMowers() \ngot = %+v,\n want %+v", got, tt.want)
			}
		})
	}
}
