package utils

type Point struct {
	X int
	Y int
}

func NewDefaultPoint() Point {
	p := Point{0, 0}
	return p
}

func (p *Point) MoveUp(count ...int) {
	num := valueOrDefault(1, count)
	p.Y -= num
}

func (p *Point) MoveDown(count ...int) {
	num := valueOrDefault(1, count)
	p.Y += num
}

func (p *Point) MoveLeft(count ...int) {
	num := valueOrDefault(1, count)
	p.X -= num
}

func (p *Point) MoveRight(count ...int) {
	num := valueOrDefault(1, count)
	p.X += num
}

func valueOrDefault(def int, count []int) int {
	if len(count) == 0 {
		return def
	} else {
		return count[0]
	}
}
