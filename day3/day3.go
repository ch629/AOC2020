package day3

import (
	"AOC2020/util"
)

type TileType byte
type Map [][]TileType

const (
	FREE TileType = iota
	TREE
)

func SolvePart1() int {
	return CountTrees(MakeMap(), 3, 1)
}

func SolvePart2() int {
	deltas := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	m := MakeMap()

	mult := 1
	for _, delta := range deltas {
		xDelta, yDelta := delta[0], delta[1]
		count := CountTrees(m, xDelta, yDelta)
		mult *= count
	}
	return mult
}

func MakeMap() Map {
	lines := util.ReadLinesPath("input.txt")
	m := Map(make([][]TileType, len(lines)))

	for y, line := range lines {
		m[y] = make([]TileType, len(line))

		for x, ch := range line {
			var tileType TileType
			if ch == '.' {
				tileType = FREE
			} else {
				tileType = TREE
			}
			m[y][x] = tileType
		}
	}

	return m
}

func CountTrees(m Map, xDelta, yDelta int) int {
	treeCount := 0
	for i, y := 0, 0; y < len(m); i, y = i + 1, y + yDelta {
		x := (i * xDelta) % len(m[y])

		if m[y][x] == TREE {
			treeCount++
		}
	}
	return treeCount
}
