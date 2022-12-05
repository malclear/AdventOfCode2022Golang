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
	stacks, linesRead := getArrayOfStacks(parsed)

	linesRead++ //skip blank line
	for i := linesRead; i < len(parsed); i++ {
		var count, fromStack, toStack = getMoveData(parsed[i])
		for j := 0; j < count; j++ {
			crate, _ := stacks[fromStack-1].Pop()
			stacks[toStack-1].Push(crate)
		}
	}
	for i := 0; i < len(stacks); i++ {
		stack, _ := stacks[i].Pop()
		fmt.Print(stack)
	}
	fmt.Println("")
	return 0
}

func getMoveData(line string) (int, int, int) {
	line = strings.ReplaceAll(line, "move ", "")
	line = strings.ReplaceAll(line, "from ", "")
	line = strings.ReplaceAll(line, "to ", "")
	line = strings.Trim(line, " ")
	values := strings.Split(line, " ")
	count, _ := strconv.Atoi(values[0])
	fromStack, _ := strconv.Atoi(values[1])
	toStack, _ := strconv.Atoi(values[2])
	return count, fromStack, toStack
}

func getArrayOfStacks(parsed []string) ([]Stack, int) {
	stacks := []Stack{
		Stack{}, Stack{}, Stack{}, Stack{}, Stack{}, Stack{}, Stack{}, Stack{}, Stack{}, Stack{},
	}
	index := 0
	for ln, containerLine := range parsed {
		index = ln
		if !strings.Contains(containerLine, "[") {
			break
		}
		containerStackList := getContainerStacks(containerLine)
		for _, crateStack := range containerStackList {
			stacks[crateStack.stack].Push(crateStack.crate)
		}
	}
	stackCount := 0
	for i, stack := range stacks {
		stackCount = i
		stack.Reverse()
		if len(stack) == 0 {
			break
		}
		stacks[i] = stack
	}
	stacks = stacks[:stackCount] //ugly hack

	return stacks, index + 1
}

func getContainerStacks(line string) []CrateAddress {

	crateAddressList := []CrateAddress{}
	splitLine := strings.Split(line, "[")
	length := 0
	for i, seg := range splitLine {
		if strings.Trim(seg, " ") == "" {
			length = len(seg)
			continue
		}

		crate := seg[0]
		stack := (length + i) / 4
		crateAddressList = append(crateAddressList, CrateAddress{crate: string(crate), stack: stack})
		length += len(seg)
	}

	return crateAddressList
}

func part2(input string) int {
	parsed := parseInput(input)
	stacks, linesRead := getArrayOfStacks(parsed)

	linesRead++ //skip blank line
	for i := linesRead; i < len(parsed); i++ {
		var count, fromStack, toStack = getMoveData(parsed[i])
		tempStack := Stack{}
		for j := 0; j < count; j++ {
			crate, _ := stacks[fromStack-1].Pop()
			tempStack.Push(crate)
		}
		for j := 0; j < count; j++ {
			crate, _ := tempStack.Pop()
			stacks[toStack-1].Push(crate)
		}
	}
	for i := 0; i < len(stacks); i++ {
		stack, _ := stacks[i].Pop()
		fmt.Print(stack)
	}
	fmt.Println("")
	return 0

}
