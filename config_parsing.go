package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var directionParsing = map[string]Direction{
	"N": North,
	"S": South,
	"E": East,
	"W": West,
}

func ParseConfigFile(content []byte) (*Lawn, *[]MowerAndCommands, error) {
	lines := strings.Split(string(content), "\n")

	topRightPos, err := parseTopRightPosition(lines[0])
	if nil != err {
		return nil, nil, err
	}

	mowers, err := parseMowers(lines[1:])

	if nil != err {
		return nil, nil, err
	}

	return topRightPos, mowers, nil
}

func parseMowers(lines []string) (*[]MowerAndCommands, error) {
	if len(lines) > 1 {

		position, direction, err := parsePositionAndDirection(lines[0])
		if nil != err {
			return nil, err
		}
		commands, err := parseCommands(lines[1])
		if nil != err {
			return nil, err
		}
		m := Mower{
			position:  *position,
			direction: *direction,
		}
		mowers, err := parseMowers(lines[2:])
		if nil != err {
			return nil, err
		}
		*mowers = append([]MowerAndCommands{{&m, commands}}, *mowers...)
		return mowers, nil
	}

	return &[]MowerAndCommands{}, nil
}

func parsePositionAndDirection(line string) (*Position, *Direction, error) {
	regexMowerPos := regexp.MustCompile("(\\d+) (\\d+) ([NSEW])")
	matchMowerPos := regexMowerPos.FindStringSubmatch(line)
	if len(matchMowerPos) == 4 {
		x, _ := strconv.Atoi(matchMowerPos[1])
		y, _ := strconv.Atoi(matchMowerPos[2])
		direction := directionParsing[matchMowerPos[3]]
		return &Position{x, y}, &direction, nil
	}
	return nil, nil, fmt.Errorf("bad format %s, must be '(\\d+) (\\d+) ([DGA])'", line)

}

func parseTopRightPosition(line string) (*Lawn, error) {
	regex := regexp.MustCompile("(\\d+) (\\d+)")
	match := regex.FindStringSubmatch(line)
	if len(match) == 3 {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		return &Lawn{Position{x, y}}, nil
	}

	return nil, fmt.Errorf("bad format %s, must be 'x y'", line)
}

func parseCommands(line string) ([]string, error) {
	regexCmd := regexp.MustCompile("[DGA]+")
	if regexCmd.MatchString(line) {
		characters := strings.Split(line, "")
		return characters, nil

	}
	return nil, fmt.Errorf("bad format %s, must be [DGA]", line)
}
