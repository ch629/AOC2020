package day2

import (
	"AOC2020/util"
	"strconv"
	"strings"
)

type Line struct {
	First     int
	Second    int
	Character rune
	Password  string
}

func MakeLine(str string) Line {
	split := strings.Split(str, ": ")
	policy, pass := split[0], split[1]

	policySplit := strings.Split(policy, " ")
	occurrences := policySplit[0]
	letter := rune(policySplit[1][0])
	occurrenceSplit := strings.Split(occurrences, "-")
	min, _ := strconv.Atoi(occurrenceSplit[0])
	max, _ := strconv.Atoi(occurrenceSplit[1])

	return Line{
		First:     min,
		Second:    max,
		Character: letter,
		Password:  pass,
	}
}

func SolveCount() int {
	lines := util.ReadLinesPath("input.txt")
	count := 0

	for _, line := range lines {
		if LinePassesCount(line) {
			count++
		}
	}

	return count
}

func SolvePlaces() int {
	lines := util.ReadLinesPath("input.txt")
	count := 0

	for _, line := range lines {
		if LinePassesPlaces(line) {
			count++
		}
	}

	return count
}

func LinePassesCount(str string) bool {
	line := MakeLine(str)

	count := 0
	for _, char := range line.Password {
		if char == line.Character {
			count++

			if count > line.Second {
				return false
			}
		}
	}

	return count >= line.First && count <= line.Second
}

func LinePassesPlaces(str string) bool {
	line := MakeLine(str)

	char1 := rune(line.Password[line.First - 1])

	var char2 rune

	if len(line.Password) < line.Second {
		char2 = ' '
	} else {
		char2 = rune(line.Password[line.Second - 1])
	}

	char1Equal := char1 == line.Character
	char2Equal := char2 == line.Character

	return (char1Equal || char2Equal) && !(char1Equal && char2Equal)
}
