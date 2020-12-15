package day1

import "AOC2020/util"

func SolvePair() int {
	ints := util.ReadNumbersPath("input.txt")
	num1, num2 := SumToPair(2020, ints)
	return num1 * num2
}

func SolveTriple() int {
	ints := util.ReadNumbersPath("input.txt")
	num1, num2, num3 := SumToTriple(2020, ints)
	return num1 * num2 * num3
}

func SumToPair(sum int, numbers []int) (int, int) {
	for i1, num1 := range numbers {
		for i2, num2 := range numbers {
			if i1 == i2 {
				continue
			}

			if num1 + num2 == sum {
				return num1, num2
			}
		}
	}

	return 0, 0
}

func SumToTriple(sum int, numbers []int) (int, int, int) {
	for i1, num1 := range numbers {
		for i2, num2 := range numbers {
			if i1 == i2 {
				continue
			}

			for i3, num3 := range numbers {
				if i1 == i3 || i2 == i3 {
					continue
				}

				if num1 + num2 + num3 == sum {
					return num1, num2, num3
				}
			}
		}
	}

	return 0, 0, 0
}
