package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
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

	var num int

	maxCalories := 0
	elfCalories := 0

	for _, str := range parsed {
		if str == "" { // trim this string var?
			if elfCalories > maxCalories {
				maxCalories = elfCalories
			}
			// this is a new Elf
			elfCalories = 0
			continue
		} else {
			num, _ = strconv.Atoi(str)
			elfCalories = elfCalories + num
		}
	}

	return maxCalories
}

func part2(input string) int {
	parsed := parseInput(input)

	var elfCalorieList []int

	var num int

	maxCalories := 0
	elfCalories := 0

	for _, str := range parsed {
		if str == "" { // trim this string var?
			if elfCalories > maxCalories {
				maxCalories = elfCalories
			}
			elfCalorieList = append(elfCalorieList, elfCalories)
			// this is a new Elf
			elfCalories = 0
			continue
		} else {
			num, _ = strconv.Atoi(str)
			elfCalories = elfCalories + num
		}
	}

	sort.Ints(elfCalorieList)
	listLen := len(elfCalorieList)
	return elfCalorieList[listLen-1] + elfCalorieList[listLen-2] + elfCalorieList[listLen-3]

}

func parseInput(input string) (ans []string) {
	for _, l := range strings.Split(input, "\n") {
		ans = append(ans, l)
	}
	return ans
}
