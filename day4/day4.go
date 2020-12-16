package day4

import (
	"AOC2020/util"
	"strings"
)

type PassportData map[string]string

// TODO: Add validation on req fields (regex/validator func)
func Solve() int {
	lines := util.ReadLinesPath("input.txt")
	count := 0
	req := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	passports := MakePassportData(lines)
	for _, data := range passports {
		if ContainsAll(data, req) {
			count++
		}
	}

	return count
}

func ContainsAll(data PassportData, req []string) bool {
	for _, r := range req {
		if _, ok := data[r]; !ok {
			return false
		}
	}
	return true
}

func MakePassportData(lines []string) []PassportData {
	joined := strings.Join(lines, " ")
	records := strings.Split(joined, "  ")
	passportData := make([]PassportData, len(records))

	for i, record := range records {
		data := make(PassportData)
		pairs := strings.Split(record, " ")

		for _, pair := range pairs {
			split := strings.Split(pair, ":")
			data[split[0]] = split[1]
		}
		passportData[i] = data
	}

	return passportData
}
