package main

import (
	"AdventOfCode2022GoLang/utils"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unsafe"
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
	fmt.Println("**************************************************************")

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
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
	OldCount           int
	DivisibleByTest    int
	ThrowToMonkeyTrue  int
	ThrowToMonkeyFalse int
	ActionValue        int
}

func part1(input string) int {
	parsed := parseInput(input)
	var monkeyList []Monkey
	var currentMonkey *Monkey

	parseMonkeyList(parsed, &monkeyList, currentMonkey)

	runMonkeyEngine(monkeyList)
	actions := []int{}
	for _, monkey := range monkeyList {
		actions = append(actions, monkey.ActionCount)
	}
	//sort.Sort(sort.IntSlice(actions))
	sort.Sort(sort.Reverse(sort.IntSlice(actions)))
	fmt.Println(parsed)
	return actions[0] * actions[1]
}

func runMonkeyEngine(monkeyList []Monkey) {
	for round := 0; round < 20; round++ {
		for monkeyIdx, _ := range monkeyList {
			for _, lootItem := range monkeyList[monkeyIdx].Loot {
				itemValue, err := strconv.Atoi(lootItem)
				if err != nil {
					fmt.Println("Error parsing loot item", monkeyIdx)
				}
				compareValue := 0
				if monkeyList[monkeyIdx].OldCount == 1 {
					compareValue = monkeyList[monkeyIdx].Operation(itemValue, monkeyList[monkeyIdx].ActionValue)
				} else {
					compareValue = monkeyList[monkeyIdx].Operation(itemValue, itemValue)
				}
				compareValue = compareValue / 3
				monkeyList[monkeyIdx].ActionCount++
				if compareValue%monkeyList[monkeyIdx].DivisibleByTest == 0 {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyTrue].Loot.Push(strconv.Itoa(compareValue))
				} else {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyFalse].Loot.Push(strconv.Itoa(compareValue))
				}
				fmt.Print()
			}
		}
	}
}

func runMonkeyEngine2(monkeyList []Monkey) {
	for round := 0; round < 1000; round++ {
		for monkeyIdx, _ := range monkeyList {
			for _, lootItem := range monkeyList[monkeyIdx].Loot {
				itemValue, err := strconv.Atoi(lootItem)
				if err != nil {
					fmt.Println("Error parsing loot item", monkeyIdx)
				}
				compareValue := 0
				if monkeyList[monkeyIdx].OldCount == 1 {
					compareValue = monkeyList[monkeyIdx].Operation(itemValue, monkeyList[monkeyIdx].ActionValue)
				} else {
					compareValue = monkeyList[monkeyIdx].Operation(itemValue, itemValue)
				}
				monkeyList[monkeyIdx].ActionCount++
				if compareValue%monkeyList[monkeyIdx].DivisibleByTest == 0 {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyTrue].Loot.Push(strconv.Itoa(compareValue))
				} else {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyFalse].Loot.Push(strconv.Itoa(compareValue))
				}
				fmt.Print()
			}
		}
	}
}

func parseMonkeyList(parsed []string, monkeyList *[]Monkey, currentMonkey *Monkey) {
	for _, line := range parsed {
		if strings.HasPrefix(line, "Monkey ") {
			// Starts with "Monkey "
			*monkeyList = append(*monkeyList, Monkey{})
			currentMonkey = &((*monkeyList)[len(*monkeyList)-1])
		}
		if strings.HasPrefix(line, "  Starting items: ") {
			// starts with " Starting items..."
			suffix := strings.TrimPrefix(line, "  Starting items: ")
			itemArray := strings.Split(suffix, " ")
			for i, s := range itemArray {
				itemArray[i] = strings.Trim(s, ",")
			}
			currentMonkey.Loot.Push(itemArray...)
		}
		if strings.HasPrefix(line, "  Operation: new = ") {
			// starts with " Operation..."
			operationText := strings.TrimPrefix(line, "  Operation: new = ")
			op, actionValue, oldCount, err := getOperation(operationText)
			if err != nil {
				fmt.Println("Error parsing operation")
				break
			}
			currentMonkey.Operation = op
			currentMonkey.ActionValue = actionValue
			currentMonkey.OldCount = oldCount
		}
		if strings.HasPrefix(line, "  Test: divisible by ") {
			divByValue, err := strconv.Atoi(strings.TrimPrefix(line, "  Test: divisible by "))
			if err != nil {
				fmt.Println("Error parsing Div By test")
				break
			}
			currentMonkey.DivisibleByTest = divByValue
		}
		if strings.HasPrefix(line, "    If true: throw to monkey ") {
			throwTrue, err := strconv.Atoi(strings.TrimPrefix(line, "    If true: throw to monkey "))
			if err != nil {
				fmt.Println("Error with 'Throw true'")
				break
			}
			currentMonkey.ThrowToMonkeyTrue = throwTrue

		}
		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			throwFalse, err := strconv.Atoi(strings.TrimPrefix(line, "    If false: throw to monkey "))
			if err != nil {
				fmt.Println("Error with 'Throw false'")
				break
			}
			currentMonkey.ThrowToMonkeyFalse = throwFalse
		}
		if strings.Trim(line, " ") == "" {
			continue
		}
	}
}

// At the end of the operation, if the pixel was passed in that operation, then light that pixel.
func part2(input string) int {
	var asdf = 0
	fmt.Println(unsafe.Sizeof(asdf))
	parsed := parseInput(input)
	var monkeyList []Monkey
	var currentMonkey *Monkey

	parseMonkeyList(parsed, &monkeyList, currentMonkey)

	runMonkeyEngine2(monkeyList)
	actions := []int{}
	for _, monkey := range monkeyList {
		actions = append(actions, monkey.ActionCount)
	}
	//sort.Sort(sort.IntSlice(actions))
	sort.Sort(sort.Reverse(sort.IntSlice(actions)))
	fmt.Println(parsed)
	return actions[0] * actions[1]
}

func getOperation(opText string) (func(int, int) int, int, int, error) {
	oldCount := strings.Count(opText, "old")
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}
	actionNumber := 0
	if oldCount == 1 {
		an, err := strconv.Atoi(re.FindString(opText))
		actionNumber = an
		if err != nil {
			fmt.Println(err)
			return nil, 0, 0, err
		}
	}

	if strings.Contains(opText, "+") {
		return addNumbers, actionNumber, oldCount, nil
	} else if strings.Contains(opText, "*") {
		return multiplyNumbers, actionNumber, oldCount, nil
	} else {
		return nil, 0, 0, err
	}
}

func addNumbers(first int, second int) int {
	return first + second
}

func multiplyNumbers(first int, second int) int {
	return first * second
}
