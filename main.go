package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

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

func part1(input string) int {
	parsed := parseInput(input)

	totalPriorities := 0
	for _, str := range parsed {
		first, second := getDividedRucksack(str)
		commonItem := getCommonItem(first, second)
		totalPriorities += getItemPriority(commonItem)
	}

	return totalPriorities
}

func part2(input string) int {
	parsed := parseInput(input)

	totalPriorities := 0
	groupPriority := 0

	for i := 0; i < len(parsed); i++ {
		if i%3 == 2 {
			groupPriority = getGroupPriority(parsed[i-2], parsed[i-1], parsed[i])
			totalPriorities += groupPriority
			groupPriority = 0
		}
	}

	return totalPriorities
}

func getGroupPriority(elf_1 string, elf_2 string, elf_3 string) int {
	common := findCommon(elf_1, elf_2)
	common = findCommon(common, elf_3)
	return getItemPriority(common)
}

func findCommon(list_1 string, list_2 string) string {
	var common string
	for _, c := range list_2 {
		if strings.Contains(list_1, string(c)) {
			common = common + string(c)
		}
	}
	return common
}

func getItemPriority(item string) int {
	asciiVal := int(item[0])
	if asciiVal >= 65 && asciiVal <= 90 {
		return asciiVal - 38 // returns the prescribed values for UPPERCASE letters
	}
	return asciiVal - 96 // returns the prescribed values for lowercase letters
}

func getCommonItem(first string, second string) string {
	var commonItem string
	for _, c := range first {
		if len(strings.Split(second, string(c))) > 1 {
			commonItem = string(c)
			break
		}
	}
	return commonItem
}

func getDividedRucksack(str string) (first string, second string) {
	first = str[0 : len(str)/2]
	second = str[len(str)/2 : len(str)]
	return first, second
}

func parseInput(input string) (ans []string) {
	for _, l := range strings.Split(input, "\n") {
		ans = append(ans, l)
	}
	return ans
}
