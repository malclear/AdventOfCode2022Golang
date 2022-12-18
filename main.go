package main

import (
	"flag"
	"fmt"
	"strconv"
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
	register_X := 1
	cycle := 1
	pointsOfInterest := map[int]int{}
	for _, line := range parsed {
		fmt.Println(line)
		if line == "noop" {
			// handle noop
			incrementCycle(&cycle, &register_X, pointsOfInterest, 0)
			continue
		}

		// handle addx
		addXValue, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			fmt.Printf("Error reading %s\n", line)
			return -1
		}
		incrementCycle(&cycle, &register_X, pointsOfInterest, 0)
		incrementCycle(&cycle, &register_X, pointsOfInterest, addXValue)
	}
	returnValue := 0
	for cycle, xValue := range pointsOfInterest {
		returnValue += cycle * xValue
	}
	return returnValue
}

func incrementCycle(cyclePnt *int, x *int, poi map[int]int, toAdd int) {
	*cyclePnt++
	*x += toAdd
	if (*cyclePnt%40)-20 == 0 {
		poi[*cyclePnt] = *x
	}
}

func incrementCycle2(cyclePnt *int, x *int, toAdd int) {
	pixel := (*cyclePnt - 1) % 40
	*cyclePnt++
	if pixel == 0 {
		fmt.Println()
	}
	if *x > pixel-2 && *x < pixel+2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	*x += toAdd
}

func part2(input string) int {
	parsed := parseInput(input)
	register_X := 1
	cycle := 1
	for _, line := range parsed {
		//fmt.Println(line)
		if line == "noop" {
			// handle noop
			incrementCycle2(&cycle, &register_X, 0)
			continue
		}

		// handle addx
		addXValue, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			fmt.Printf("Error reading %s\n", line)
			return -1
		}
		incrementCycle2(&cycle, &register_X, 0)
		incrementCycle2(&cycle, &register_X, addXValue)
	}
	fmt.Println()
	return -1
}
