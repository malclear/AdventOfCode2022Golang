package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
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

func part1(input string) int {
	parsed := parseInput(input)
	indicesTotal := 0
	for i := 0; i < (len(parsed)+1)/3; i++ {
		//left := Expression{Encoding: parsed[i*3]}
		fmt.Println(parsed[i*3])
		fmt.Println(parsed[i*3+1])
		left, _, _ := parse(tokenize(parsed[i*3]))
		right, _, _ := parse(tokenize(parsed[i*3+1]))

		if left.Compare(right) <= 0 {
			indicesTotal += i + 1
			fmt.Println("Pair", i+1, "is ordered")
		}
		fmt.Println()
	}
	return indicesTotal
}

func part2(input string) int {
	parsed := parseInput(input)
	expressions := ExpressionSet{}
	for _, line := range parsed {
		if strings.Trim(line, " ") == "" {
			continue
		}
		expression, _, _ := parse(tokenize(line))
		expressions = append(expressions, expression)
	}

	alpha, _, _ := parse(tokenize("[[2]]"))
	beta, _, _ := parse(tokenize("[[6]]"))
	expressions = append(expressions, alpha)
	expressions = append(expressions, beta)

	expressions.Swap(len(expressions)-1, len(expressions)-2)
	var alphaIndex, betaIndex int
	sort.Sort(expressions)
	for i, _ := range expressions {
		if (expressions[i]) == alpha {
			alphaIndex = i + 1
		}
		if (expressions[i]) == beta {
			betaIndex = i + 1
		}
	}
	return alphaIndex * betaIndex
}

type ExpressionSet []*Expression

func (p ExpressionSet) Less(i, j int) bool {
	return (p[i]).Compare(p[j]) < 0
}

func (p ExpressionSet) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ExpressionSet) Len() int {
	return len(p)
}

type Expression struct {
	Expressions []*Expression
	Value       *int
}

func tokenize(string string) []string {
	tokenRegexp := regexp.MustCompile(`\[|\]|[\d]+`)
	return tokenRegexp.FindAllString(string, -1)
}

func parse(tokens []string) (*Expression, int, error) {
	// Base case: an empty list of tokens represents an empty list
	if len(tokens) == 0 {
		return &Expression{}, 0, nil
	}
	// If the first token is an opening bracket, we have a list
	if tokens[0] == "[" {
		// Get a new instance of Expression and initialize its child list to an empty array
		result := &(Expression{Expressions: []*Expression{}})
		// Consume tokens until we reach the closing bracket
		i := 1
		for tokens[i] != "]" {
			// Recursively parse the tokens inside the list
			pt, innerTokenCount, err := parse(tokens[i:])
			if err != nil {
				fmt.Println("WHAAAAAT!")
			}
			result.Expressions = append(result.Expressions, pt)
			// Skip over the tokens we just parsed, plus the closing bracket
			i += innerTokenCount
		}
		// add 1 to i to include the closing bracket
		return result, i + 1, nil
	}
	// If the first token is not an opening bracket, it must be an atom
	// Try to convert it to an integer
	if intVal, err := strconv.Atoi(tokens[0]); err == nil {
		return &(Expression{Value: &intVal}), 1, nil
	}

	// If the first token is not an opening bracket, it must be an integer
	// If it's not an integer, return an error
	return nil, -1, fmt.Errorf("invalid token: %s", tokens[0])
}

func (e *Expression) Compare(other *Expression) int {
	if (e.Value != nil && e.Expressions != nil) || (other.Value != nil && other.Expressions != nil) {
		panic("Can't have both be non-nil!! ")
	}

	if e.Value != nil && other.Value != nil {
		return *(e.Value) - *(other.Value)
	}

	if e.Expressions != nil && other.Expressions != nil {
		if len(e.Expressions)+len(other.Expressions) == 0 {
			return 0
		}
		if len(e.Expressions) > 0 && len(other.Expressions) > 0 {
			i := 0
			for len(e.Expressions) > i && len(other.Expressions) > i {
				r := e.Expressions[i].Compare(other.Expressions[i])
				if r == 0 {
					i++
					continue
				}
				return r
			}
			return len(e.Expressions) - len(other.Expressions)
		} else {
			return len(e.Expressions) - len(other.Expressions)
		}
	}

	// Elevate the lone integer to be part of a list, as its comparator.
	left := e
	if e.Value != nil {
		left = &(Expression{Expressions: []*Expression{e}})
	}

	right := other
	if other.Value != nil {
		right = &(Expression{Expressions: []*Expression{other}})
	}

	return left.Compare(right)
}
