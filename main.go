package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"regexp"
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

func part1(input string) int {
	parsed := parseInput(input)

	orderedPairCount := 0
	pairs := getPairs(parsed)
	for _, pair := range pairs {
		//if pair.Left.Compare(pair.Right) <= 0 {
		//	orderedPairCount++
		//}
		fmt.Println(pair)
	}
	return orderedPairCount
}

func getPairs(parsed []string) map[int]Pair {
	pairs := make(map[int]Pair)

	for i := 0; i < (len(parsed)+1)/3; i++ {
		//left := Expression{Encoding: parsed[i*3]}

		thing := parse(tokenize(parsed[i*3]))
		fmt.Println(thing)

		//left.Init()
		//right := Expression{Encoding: parsed[i*3+1]}
		//right.Init()
		//pair := Pair{Left: left, Right: right}
		//pairs[i] = pair
	}
	return pairs
}

// At the end of the operation, if the pixel was passed in that operation, then light that pixel.
func part2(input string) int {
	parsed := parseInput(input)
	fmt.Println(parsed)
	return -1
}

type Pair struct {
	Left  Expression
	Right Expression
}

type Expression struct {
	Encoding    string
	Expressions []*Expression
	Value       *int
}

//
//func  NewExpression(encString string) *Expression {
//	/*
//		1) Assert the "Encoding" string starts with a "[" or a numeric digit.
//		2) If "[", find all text before the ending "]" and
//			create a new Expression instance with the enclosed string.
//			2a) Call Init on the new instance
//			2b) add the new instance to the list this instance
//		3) If the first char is a numeric, split the string on commas and trim the results.
//			3a) For each integer, create a new expression and assign the value to the "Value" property
//	*/
//	retval := Expression{Encoding: encString}
//	if
//	var createExpressionSet bool
//	var createIntegerExpression bool
//	if strings.HasPrefix(e.Encoding, "[") {
//		createExpressionSet = true
//	}
//
//}

func tokenize(string string) []string {
	tokenRegexp := regexp.MustCompile(`\[|\]|[\d]+`)
	return tokenRegexp.FindAllString(string, -1)
}

func parse(tokens []string) interface{} {
	// Base case: an empty list of tokens represents an empty list
	if len(tokens) == 0 {
		return []interface{}{}
	}
	// If the first token is an opening bracket, we have a list
	if tokens[0] == "[" {
		result := []interface{}{}
		// Consume tokens until we reach the closing bracket
		i := 1
		for tokens[i] != "]" {
			// Recursively parse the tokens inside the list
			result = append(result, parse(tokens[i:]))
			// Skip over the tokens we just parsed, plus the closing bracket
			i += len(result[len(result)-1].([]interface{})) + 1
		}
		return result
	}
	// If the first token is not an opening bracket, it must be an atom
	// Try to convert it to an integer
	if intVal, err := strconv.Atoi(tokens[0]); err == nil {
		return intVal
	}
	// If it's not an integer, return an error
	return fmt.Errorf("invalid token: %s", tokens[0])
}

func (e *Expression) Compare(other *Expression) int {

	return 0
}

func (e *Expression) StartBracketedExpression() *Expression { return nil }
func (e *Expression) EndBracketedExpression() *Expression   { return nil }
func (e *Expression) CreateIntegerExpression() *Expression  { return nil }

//func Tokenize(encStr string) []int {
//	for _, c := range encStr {
//		fmt.Println(c, string(c))
//	}
//	return []int{0}
//}

const (
	// "["
	OPEN_BRACKET int = 91
	// "]"
	CLOSED_BRACKET = 93
	// "0"
	ZERO = 48
	// "1"
	ONE = 49
	// "2"
	TWO = 50
	// "3"
	THREE = 51
	// "4"
	FOUR = 52
	// "5"
	FIVE = 53
	// "6"
	SIX = 54
	// "7"
	SEVEN = 55
	// "8"
	EIGHT = 56
	// "9"
	NINE = 57
	// ","
	COMMA = 44
	// " "
	SPACE = 32
)
