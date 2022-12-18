package main

import (
	"flag"
	"fmt"
	"math"
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
	head := Point{0, 0}
	tail := Point{0, 0}
	var tailPositions map[Point]struct{}
	tailPositions = make(map[Point]struct{})
	fmt.Println(tailPositions)

	for _, moveset := range parsed {
		moves := strings.Split(moveset, " ")
		direction := moves[0]
		count, _ := strconv.ParseInt(moves[1], 10, 32)

		for i := 0; i < int(count); i++ {
			previousHead := head
			switch direction {
			case "U":
				head.MoveUp()
			case "D":
				head.MoveDown()
			case "L":
				head.MoveLeft()
			default:
				head.MoveRight()
			}
			fmt.Print(head)
			fmt.Print(" --> ")
			tail = AdjustNext(tail, head, previousHead)
			tailPositions[tail] = struct{}{}
			fmt.Println(tail)
		}
	}

	return len(tailPositions)
}

func part2(input string) int {
	parsed := parseInput(input)
	nodes := []Point{
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
	}

	var uniqueTailPositions map[Point]struct{}
	uniqueTailPositions = make(map[Point]struct{})

	//for _, moveset := range parsed {
	for _, moveset := range parsed {

		direction := strings.Split(moveset, " ")[0]
		moveCount, _ := strconv.ParseInt(strings.Split(moveset, " ")[1], 10, 32)

		for i := 0; i < int(moveCount); i++ {
			previousHead := nodes[0]
			node := nodes[0]
			switch direction {
			case "U":
				node.MoveUp()
			case "D":
				node.MoveDown()
			case "L":
				node.MoveLeft()
			default:
				node.MoveRight()
			}
			nodes[0] = node
			for nodeIdx := 1; nodeIdx < 10; nodeIdx++ {
				followerPreviousNode := nodes[nodeIdx]
				nodes[nodeIdx] = AdjustNext(nodes[nodeIdx], nodes[nodeIdx-1], previousHead)
				previousHead = followerPreviousNode

			}
			uniqueTailPositions[nodes[9]] = struct{}{}

			//fmt.Print(head)
			//fmt.Print(" --> ")
			//tail = AdjustNext(tail, head, previousHead)
		}

		printGraph(nodes, moveset)
		fmt.Println("")
	}

	return len(uniqueTailPositions)

}

func printGraph(nodes []Point, moveset string) {
	fmt.Printf("%s\n", moveset)

	for row := -35; row < 26; row++ {
		for col := -31; col < 34; col++ {
			space := "."
			for n := 0; n < 10; n++ {
				if nodes[n].X == col && nodes[n].Y == row {
					space = strconv.Itoa(n)
					if n == 0 {
						space = "H"
					}
					break
				}
			}
			if row == 0 && col == 0 {
				space = "s"
			}
			fmt.Print(space + " ")
		}

		fmt.Println()
	}
	fmt.Println()
}

func AdjustNext(next Point, leader Point, previousLeaderPos Point) Point {
	distance := math.Sqrt(float64(((leader.X - next.X) * (leader.X - next.X)) + ((leader.Y - next.Y) * (leader.Y - next.Y))))
	touching := distance < 2
	newT := next
	//if leader.X != next.X && leader.Y != next.Y && !touching {
	if distance > 2 {
		if leader.X > newT.X {
			newT.MoveRight()
		} else {
			newT.MoveLeft()
		}
		if leader.Y > newT.Y {
			newT.MoveDown()
		} else {
			newT.MoveUp()
		}
	} else if distance == 2 {
		if leader.X > newT.X {
			newT.MoveRight()
		}
		if leader.Y > newT.Y {
			newT.MoveDown()
		}
		if leader.X < newT.X {
			newT.MoveLeft()
		}
		if leader.Y < newT.Y {
			newT.MoveUp()
		}
	} else {
		if !touching {
			newT = previousLeaderPos
		}
	}
	return newT
}

func printUniquePostions(positions map[Point]struct{}) {
	hightestX := 0
	lowestX := 0
	hightestY := 0
	lowestY := 0
	for k, _ := range positions {
		if k.X > hightestX {
			hightestX = k.X
		}
		if k.X < lowestX {
			lowestX = k.X
		}
		if k.Y > hightestY {
			hightestY = k.Y
		}
		if k.Y < lowestY {
			lowestY = k.Y
		}
	}

}

func CheckViewScoreForPoint(point Point, forest Forest, xMax int, yMax int) int {
	treeHeight := forest.Trees[point].Height

	countDown := 0
	index := point
	index.MoveDown()
	for index.Y < forest.Height {
		countDown++
		if forest.Trees[index].Height >= treeHeight {
			break
		}

		index.MoveDown()
	}

	countUp := 0
	index = point
	index.MoveUp()
	for index.Y >= 0 {
		countUp++
		if forest.Trees[index].Height >= treeHeight {
			break
		}

		index.MoveUp()
	}

	countRight := 0
	index = point
	index.MoveRight()
	for index.X < forest.Width {
		countRight++
		if forest.Trees[index].Height >= treeHeight {
			break
		}

		index.MoveRight()
	}

	countLeft := 0
	index = point
	index.MoveLeft()
	for index.X >= 0 {
		countLeft++
		if forest.Trees[index].Height >= treeHeight {
			break
		}

		index.MoveLeft()
	}

	return countUp * countDown * countLeft * countRight
}

func GetForestWithVisibleTrees(parsed []string) Forest {
	var forest Forest
	//forest.Trees = make(map[Point]Tree)
	forest.Init(len(parsed[0]), len(parsed))
	for y, line := range parsed {
		for x := 0; x < len(line); x++ {
			size := int(line[x]) - 48
			point := Point{x, y}

			tree := Tree{size, 15}
			forest.Trees[point] = tree
		}
	}
	var tallest int
	for y := 0; y < len(parsed); y++ {
		tallest = 0
		for x := 0; x < len(parsed[y]); x++ {
			f := forest.Trees[Point{x, y}]
			if f.Height > tallest {
				tallest = forest.Trees[Point{x, y}].Height
			} else {
				f.IsVisible = Toggle(f.IsVisible, LEFT)
			}
			forest.Trees[Point{x, y}] = f
		}
	}

	for y := 0; y < len(parsed); y++ {
		tallest = 0
		for x := len(parsed[y]) - 1; x > 0; x-- {
			f := forest.Trees[Point{x, y}]
			if f.Height > tallest {
				tallest = forest.Trees[Point{x, y}].Height
			} else {
				f.IsVisible = Toggle(f.IsVisible, RIGHT)
			}
			forest.Trees[Point{x, y}] = f
		}
	}

	for x := 0; x < len(parsed[0]); x++ {
		tallest = 0
		for y := len(parsed) - 1; y > 0; y-- {
			f := forest.Trees[Point{x, y}]
			if f.Height > tallest {
				tallest = forest.Trees[Point{x, y}].Height
			} else {
				f.IsVisible = Toggle(f.IsVisible, UP)
			}
			forest.Trees[Point{x, y}] = f
		}
	}
	for x := 0; x < len(parsed[0]); x++ {
		tallest = 0
		for y := 0; y < len(parsed); y++ {
			f := forest.Trees[Point{x, y}]
			if f.Height > tallest {
				tallest = forest.Trees[Point{x, y}].Height
			} else {
				f.IsVisible = Toggle(f.IsVisible, DOWN)
			}
			forest.Trees[Point{x, y}] = f
		}
	}

	return forest
}
