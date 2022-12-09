package main

type Forest struct {
	Trees  map[Point]Tree
	Height int
	Width  int
}

func (f *Forest) Init(width int, height int) {
	f.Height = height
	f.Width = width
	f.Trees = make(map[Point]Tree)
}
