package util

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func ReadNumbersPath(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return ReadNumbers(f)
}

func ReadLinesPath(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return ReadLines(f)
}

func ReadNumbers(file io.ReadCloser) []int {
	lines := ReadLines(file)
	ints := make([]int, len(lines))

	for i, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		ints[i] = num
	}
	return ints
}

func ReadLines(file io.ReadCloser) []string {
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
