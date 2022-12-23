package main

import (
	"AdventOfCode2022GoLang/utils"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
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

var mapHeight = 0
var mapWidth = 0
var puzzleMap = make(map[utils.Point]*MapCell)
var pathHistory utils.Stack
var startingPoint, endPoint *MapCell
var bestPath utils.Stack
var depth = 0

const (
	NORTH = 1 << iota
	EAST
	SOUTH
	WEST
)

func IsLegalDirection(direction int, thisPoint MapCell) bool {
	current := thisPoint.Location
	var newPoint utils.Point
	if direction == NORTH {
		newPoint = utils.Point{current.X, current.Y - 1}
	} else if direction == EAST {
		newPoint = utils.Point{current.X + 1, current.Y}
	} else if direction == SOUTH {
		newPoint = utils.Point{current.X, current.Y + 1}
	} else {
		newPoint = utils.Point{current.X - 1, current.Y}
	}
	if !IsOnMap(newPoint) {
		return false
	}

	if puzzleMap[newPoint].Height <= thisPoint.Height+1 {
		return true
	}
	return false
}

func IsOnMap(point utils.Point) bool {
	if point.X < 0 || point.Y < 0 {
		return false
	}
	if point.X >= mapWidth || point.Y >= mapHeight {
		return false
	}
	return true
}

type MapCell struct {
	Location         utils.Point
	LegalTransitions int
	Height           rune
	MaxSteps         int
	Visited          bool
}

func part1(input string) int {
	parsed := parseInput(input)
	mapWidth = len(parsed[0])
	mapHeight = len(parsed)
	for y, line := range parsed {
		for x, cell := range line {
			point := utils.Point{X: x, Y: y}
			puzzleMap[point] = &(MapCell{Height: cell, Location: point, MaxSteps: mapWidth * mapHeight})
			if string(cell) == "S" {
				puzzleMap[point].Height = rune("a"[0])
				startingPoint = puzzleMap[point]
			}
			if string(cell) == "E" {
				puzzleMap[point].Height = rune("z"[0])
				endPoint = puzzleMap[point]
			}
		}
	}
	for _, mapCell := range puzzleMap {
		if IsLegalDirection(NORTH, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | NORTH
		}
		if IsLegalDirection(EAST, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | EAST
		}
		if IsLegalDirection(SOUTH, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | SOUTH
		}
		if IsLegalDirection(WEST, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | WEST
		}
	}

	for direction := 1; direction < 16; direction *= 2 {
		startingPoint.TryMove(int(direction))
	}

	return bestPath.Size()
}

func (fromPoint *MapCell) TryMove(direction int) {
	if (fromPoint.LegalTransitions & direction) == 0 {
		return
	}

	newLocation := fromPoint.Location
	newLocation.MoveDirection(direction)
	newMapCell := puzzleMap[newLocation]
	if newMapCell.MaxSteps < depth+1 || (newMapCell.MaxSteps == depth+1 && newMapCell.Visited) {
		return
	}
	if pathHistory.Contains(newLocation) {
		return
	}
	if pathHistory.Size()+1 >= bestPath.Size() && (bestPath.Size() > 0) {
		return
	}

	fmt.Printf("X: %d, Y: %d, MaxSteps: %d, depth: %d\n", newLocation.X+1, newLocation.Y+1, newMapCell.MaxSteps, depth+1)

	// push to pathHistory
	pathHistory.Push(newMapCell.Location)
	depth++
	newMapCell.MaxSteps = depth
	newMapCell.Visited = true

	// if this is the destination,
	//		save to BestPath
	if newMapCell.Location == endPoint.Location {
		if bestPath.Size() == 0 || pathHistory.Size() < bestPath.Size() {
			bestPath = pathHistory
			fmt.Println(bestPath.Size())
		}
	}

	// for each newDirection,
	// 		call TryMove() from here
	for newDirection := 1; newDirection < 16; newDirection *= 2 {
		if newDirection+direction == 5 || direction+newDirection == 10 {
			continue
		}
		newMapCell.TryMove(int(newDirection))
	}

	// pop from history
	pathHistory.Pop()
	depth--
	// return
}

func findBestPath(terrainMap map[int]map[int]*MapCell) int {
	terrainMap[0][0].Height = rune(100)
	return 1
}

func part2(input string) int {
	parsed := parseInput(input)
	mapWidth = len(parsed[0])
	mapHeight = len(parsed)
	for y, line := range parsed {
		for x, cell := range line {
			point := utils.Point{X: x, Y: y}
			puzzleMap[point] = &(MapCell{Height: cell, Location: point, MaxSteps: mapWidth * mapHeight})
			if string(cell) == "S" {
				puzzleMap[point].Height = rune("a"[0])
			}
			if string(cell) == "E" {
				puzzleMap[point].Height = rune("z"[0])
				endPoint = puzzleMap[point]
			}
		}
	}
	for _, mapCell := range puzzleMap {
		if IsLegalDirection(NORTH, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | NORTH
		}
		if IsLegalDirection(EAST, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | EAST
		}
		if IsLegalDirection(SOUTH, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | SOUTH
		}
		if IsLegalDirection(WEST, *mapCell) {
			mapCell.LegalTransitions = mapCell.LegalTransitions | WEST
		}
	}
	for _, cell := range puzzleMap {
		if cell.Height == rune("a"[0]) {
			startingPoint = cell
			for direction := 1; direction < 16; direction *= 2 {
				startingPoint.TryMove(int(direction))
			}
		}
	}
	//return bestPath.Size()
	return bestPath.Size()
}
