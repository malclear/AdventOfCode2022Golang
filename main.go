package main

import (
	"AdventOfCode2022GoLang/utils"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"math/big"
	"regexp"
	"sort"
	"strconv"
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
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	fmt.Println("**************************************************************")

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
	var monkeyList []Monkey
	var currentMonkey *Monkey

	parseMonkeyList(parsed, &monkeyList, currentMonkey, 1)

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
	productOfAllTests := 1
	for _, m := range monkeyList {
		productOfAllTests *= m.DivisibleByTest
	}
	for round := 0; round < 10000; round++ {
		for monkeyIdx, _ := range monkeyList {
			for _, lootItem := range monkeyList[monkeyIdx].Loot {
				oldValue, err := strconv.Atoi(lootItem)
				if err != nil {
					fmt.Println("Error parsing loot item", monkeyIdx)
					break
				}
				newValue := 0
				if monkeyList[monkeyIdx].OldCount == 1 {
					newValue = monkeyList[monkeyIdx].Operation(oldValue, monkeyList[monkeyIdx].ActionValue)
				} else {
					newValue = monkeyList[monkeyIdx].Operation(oldValue, oldValue)
				}
				//newValue = newValue / 3
				//newValue = (newValue % monkeyList[monkeyIdx].DivisibleByTest) + monkeyList[monkeyIdx].DivisibleByTest
				newValue = newValue % productOfAllTests
				fmt.Println(newValue)
				monkeyList[monkeyIdx].ActionCount++
				if newValue%monkeyList[monkeyIdx].DivisibleByTest == 0 {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyTrue].Loot.Push(strconv.Itoa(newValue))
				} else {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyFalse].Loot.Push(strconv.Itoa(newValue))
				}
				fmt.Print()
			}
		}
	}
}

func runMonkeyEngine2(monkeyList []Monkey) {

	for round := 0; round < 1000; round++ {
		fmt.Println("Starting:", round)
		for monkeyIdx, _ := range monkeyList {
			for _, lootItem := range monkeyList[monkeyIdx].Loot {
				//itemValue, err := strconv.Atoi(lootItem)
				//if err != nil {
				//	fmt.Println("Error parsing loot item", monkeyIdx)
				//}
				compareValue := big.NewInt(0)
				itemValue := big.NewInt(0)
				itemValue.SetString(lootItem, 10)
				actionValue := big.NewInt(0)
				actionValue.SetInt64(int64(monkeyList[monkeyIdx].ActionValue))

				if monkeyList[monkeyIdx].OldCount == 1 {
					compareValue = monkeyList[monkeyIdx].OperationBig(itemValue, actionValue)
				} else {
					compareValue = monkeyList[monkeyIdx].OperationBig(itemValue, itemValue)
				}

				monkeyList[monkeyIdx].ActionCount++
				modulo := big.NewInt(0)
				divByTest := big.NewInt(0)
				divByTest.SetInt64(int64(monkeyList[monkeyIdx].DivisibleByTest))
				zero := big.NewInt(0)
				zero.SetInt64(0)
				if modulo.Mod(compareValue, divByTest).Cmp(zero) == 0 {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyTrue].Loot.Push(compareValue.String())
				} else {
					monkeyList[monkeyIdx].Loot.Pop()
					monkeyList[monkeyList[monkeyIdx].ThrowToMonkeyFalse].Loot.Push(compareValue.String())
				}
				fmt.Print()
			}
		}
	}
}

func parseMonkeyList(parsed []string, monkeyList *[]Monkey, currentMonkey *Monkey, partNumber int) {
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
			if partNumber == 1 {
				op, actionValue, oldCount, err := getOperation(operationText)
				if err != nil {
					fmt.Println("Error parsing operation")
					break
				}
				currentMonkey.Operation = op
				currentMonkey.ActionValue = actionValue
				currentMonkey.OldCount = oldCount
			} else {
				op, actionValue, oldCount, err := getBigOperation(operationText)
				if err != nil {
					fmt.Println("Error parsing operation")
					break
				}
				currentMonkey.OperationBig = op
				currentMonkey.ActionValue = actionValue
				currentMonkey.OldCount = oldCount
			}
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
	parsed := parseInput(input)
	var monkeyList []Monkey
	var currentMonkey *Monkey

	parseMonkeyList(parsed, &monkeyList, currentMonkey, 1)

	runMonkeyEngine(monkeyList)
	actions := []int{}
	for _, monkey := range monkeyList {
		actions = append(actions, monkey.ActionCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(actions)))
	return actions[0] * actions[1]
}

func getBigOperation(opText string) (func(*big.Int, *big.Int) *big.Int, int, int, error) {
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
		return addBigNumbers, actionNumber, oldCount, nil
	} else if strings.Contains(opText, "*") {
		return multiplyBigNumbers, actionNumber, oldCount, nil
	} else {
		return nil, 0, 0, err
	}
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

func addBigNumbers(first *big.Int, second *big.Int) *big.Int {
	sum := big.NewInt(0)
	sum = sum.Add(first, second)
	return sum
}

func multiplyBigNumbers(first *big.Int, second *big.Int) *big.Int {
	product := big.NewInt(0)
	product = product.Mul(first, second)
	//fmt.Println(first, " * ", second, " = ", product)
	return product
}
