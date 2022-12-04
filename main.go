package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/cast"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func parseInput(input string) (ans []string) {
	for _, l := range strings.Split(input, "\n") {
		ans = append(ans, l)
	}
	return ans
}

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	countOfFullyContainedSegments := 0

	for _, line := range parsed {
		elfRanges := strings.Split(line, ",")
		lo1, hi1 := getHiLoOfRange(elfRanges[0])
		lo2, hi2 := getHiLoOfRange(elfRanges[1])
		if (lo1 <= lo2 && hi1 >= hi2) || (lo2 <= lo1 && hi2 >= hi1) {
			countOfFullyContainedSegments++
			fmt.Println(line)
		}
	}
	return countOfFullyContainedSegments
}

func part2(input string) int {
	parsed := parseInput(input)
	partiallyContainedSegments := 0

	for _, line := range parsed {
		elfRanges := strings.Split(line, ",")
		lo1, hi1 := getHiLoOfRange(elfRanges[0])
		lo2, hi2 := getHiLoOfRange(elfRanges[1])
		if (lo1 <= lo2 && hi1 >= lo2) || (lo2 <= lo1 && hi2 >= lo1) ||
			(lo1 >= lo2 && hi1 <= lo2) || (lo2 >= lo1 && hi2 <= lo1) {
			partiallyContainedSegments++
			fmt.Println(line)
		}
	}
	return partiallyContainedSegments
}

func getHiLoOfRange(rangeUnits string) (int, int) {
	lo := strings.Split(rangeUnits, "-")[0]
	hi := strings.Split(rangeUnits, "-")[1]
	return cast.ToInt(lo), cast.ToInt(hi)
}
