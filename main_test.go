package main

import (
	"AdventOfCode2022GoLang/utils"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"strings"
	"testing"
)

//   func TestFoo(t *testing.T) {
//       // <setup code>
//       t.Run("A=1", func(t *testing.T) { ... })
//       t.Run("A=2", func(t *testing.T) { ... })
//       t.Run("B=1", func(t *testing.T) { ... })
//       // <tear-down code>
//   }

//func Test_Bits(t *testing.T) {
//	var b utils.Bit
//	b.Set(utils.UP)
//	b.Toggle(utils.DOWN)
//
//	got := b.Has(utils.UP)
//	if !got {
//		t.Errorf("Setting UP = false, wanted true")
//	}
//
//	got = b.Has(utils.DOWN)
//	if !got {
//		t.Errorf("Toggling DOWN = false, wanted true")
//	}
//
//	got = b.Has(utils.LEFT)
//	if got {
//		t.Errorf("initial LEFT = true, wanted false")
//	}
//}

//func Test_Bits2(t *testing.T) {
//	var b utils.Bit
//	b.Set(utils.UP).Toggle(utils.DOWN)
//
//	got := b.Has(utils.UP)
//	if !got {
//		t.Errorf("Setting UP = false, wanted true")
//	}
//
//	got = b.Has(utils.DOWN)
//	if !got {
//		t.Errorf("Toggling DOWN = false, wanted true")
//	}
//
//	got = b.Has(utils.LEFT)
//	if got {
//		t.Errorf("initial LEFT = true, wanted false")
//	}
//
//}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", util.ReadFile("input2.txt"), 2400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 1789},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Queue(t *testing.T) {
	var q utils.Queue
	q.Push("first")
	q.Push("second")
	q.Push("third")
	if len(q) != 3 {
		t.Error("Queue was incorrect size")
	}
	asdf, success := q.Pop()
	if !success {
		t.Error("Could not pop appropriately")
	}
	if asdf != "first" {
		t.Error("Did not pop 'first', as expected")
	}
	if len(q) != 2 {
		t.Error("Queue was incorrect size after pop")
	}
}

func Test_Queue2(t *testing.T) {
	var q utils.Queue
	q.Push("first", "second", "third")

	if len(q) != 3 {
		t.Error("Queue was incorrect size")
	}
	item, success := q.Pop()
	if !success {
		t.Error("Could not pop appropriately")
	}
	if item != "first" {
		t.Error("Did not pop 'first', as expected")
	}
	if len(q) != 2 {
		t.Error("Queue was incorrect size after pop")
	}
	item, _ = q.Pop()
	if item != "second" {
		t.Error("Right pop didn't work")
	}
}

func Test_StringParsing(t *testing.T) {
	asdf := "ASDF kdkdkd"
	stripped := ""
	if strings.HasPrefix(asdf, "ASDF ") {
		stripped = strings.TrimPrefix(asdf, "ASDF ")
	}
	fmt.Print(stripped)
}
