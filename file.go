package main

type File struct {
	Name string
	Size int
}

func (f *File) Init(name string, size int) {
	f.Name = name
	f.Size = size
}
