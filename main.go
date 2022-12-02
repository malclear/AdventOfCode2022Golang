package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

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

	score := 0

	var theirs, mine string
	for _, str := range parsed {
		theirs = strings.Split(str, " ")[0]
		mine = strings.Split(str, " ")[1]

		result := roundResult(theirs, mine)
		score = score + getRoundScore(mine, result)
	}

	return score

}

func part2(input string) int {
	parsed := parseInput(input)

	score := 0

	var theirs string
	for _, str := range parsed {
		theirs = strings.Split(str, " ")[0]

		// the following converts either X, Y, or Z into -1, 0, 1 as a result.
		result := int(([]rune(strings.Split(str, " ")[1]))[0]) - 89

		mine := getProperThrow(theirs, result)
		score = score + getRoundScore(mine, result)
	}

	return score
}

func getProperThrow(theirs string, result int) string {
	myPotentialThrows := []string{"X", "Y", "Z"}

	idx := (int(([]rune(theirs))[0]) - 65 + result + 3) % 3
	return myPotentialThrows[idx]

}

func roundResult(theirs string, mine string) int {
	if theirs == "A" { // ROCK
		if mine == "X" {
			return 0
		}
		if mine == "Y" {
			return 1
		}
		if mine == "Z" {
			return -1
		}
	}

	if theirs == "B" { // PAPER
		if mine == "X" {
			return -1
		}
		if mine == "Y" {
			return 0
		}
		if mine == "Z" {
			return 1
		}

	}
	if theirs == "C" { // Scissors
		if mine == "X" {
			return 1
		}
		if mine == "Y" {
			return -1
		}
		if mine == "Z" {
			return 0
		}
	}
	return 0
}

func getRoundScore(throw string, result int) int {
	if throw == "X" || throw == "A" {
		return 1 + 3*(result+1)
	}
	if throw == "Y" || throw == "B" {
		return 2 + 3*(result+1)
	}
	if throw == "Z" || throw == "C" {
		return 3 + 3*(result+1)
	}

	return 0
}

func parseInput(input string) (ans []string) {
	for _, l := range strings.Split(input, "\n") {
		ans = append(ans, l)
	}
	return ans
}
