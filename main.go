package main

import (
	"flag"
	"fmt"
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
		ans := part1(util.ReadFile("./inputSample.txt"))
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
	forest := GetForestWithVisibleTrees(parsed)

	treeCount := 0
	for y := 0; y < len(parsed); y++ {
		for x := 0; x < len(parsed[0]); x++ {
			tree := forest.Trees[Point{x, y}]
			if tree.IsVisible > 0 {
				fmt.Print(string(byte(tree.Height)))
				treeCount++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

	fmt.Print("Total tree count is ")
	fmt.Println(treeCount)

	return 0
}

func part2(input string) int {
	parsed := parseInput(input)
	forest := GetForestWithVisibleTrees(parsed)

	bestView := 0
	var pointWithBestView Point
	for point, _ := range forest.Trees {
		view := CheckViewScoreForPoint(point, forest, len(parsed[0]), len(parsed))
		if view > bestView {
			pointWithBestView = point
			bestView = view
		}
	}

	fmt.Println(pointWithBestView)
	return bestView

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
