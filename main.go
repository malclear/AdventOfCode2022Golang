package main

import (
	"AdventOfCode2022GoLang/utils"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"math/big"
	"strings"
)

func parseInput(input string) (ans []string) {
	for _, l := range strings.Split(input, "\n") {
		ans = append(ans, l)
	}
	return ans
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	fmt.Println("**************************************************************")

	if part == 1 {
		ans := part1(util.ReadFile("./inputSample.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./inputSample.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

type Monkey struct {
	Loot               utils.Queue
	ActionCount        int
	Operation          func(int, int) int
	OperationBig       func(*big.Int, *big.Int) *big.Int
	OldCount           int
	DivisibleByTest    int
	ThrowToMonkeyTrue  int
	ThrowToMonkeyFalse int
	ActionValue        int
}

func part1(input string) int {
	parsed := parseInput(input)
	fmt.Println(parsed)

	return -1
}

// At the end of the operation, if the pixel was passed in that operation, then light that pixel.
func part2(input string) int {
	parsed := parseInput(input)
	fmt.Println(parsed)
	return -1
}
