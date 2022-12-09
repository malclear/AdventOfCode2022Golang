package main

type Tree struct {
	Height    int
	IsVisible Bit
}

func (t *Tree) Init(height int) {
	t.Height = height
}
