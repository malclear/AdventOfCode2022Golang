package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	Name             string
	Path             string
	ChildDirectories []*Directory
	Files            []*File
}

func NewDirectoryByName(name string, parent *Directory) *Directory {
	var pwd string
	if parent == nil {
		name = ""
		pwd = ""
		return NewDirectoryByPath("/")
	}

	if parent.Path == "/" {
		pwd = ""
	} else {
		pwd = parent.Path
	}
	return NewDirectoryByPath(pwd + "/" + name)
}

func NewDirectoryByPath(fullPath string) *Directory {
	splitPath := strings.Split(fullPath, "/")
	newName := splitPath[len(splitPath)-1]
	var d Directory
	d.Name = newName
	d.Path = strings.Join(splitPath, "/")
	d.ChildDirectories = []*Directory{}
	d.Files = []*File{}
	return &d
}

func (d *Directory) AddFiles(files []*File) {
	d.Files = files
}

func (d *Directory) AddDirectories(directories []*Directory) {
	d.ChildDirectories = directories
}

func (d *Directory) GetTotalSize() int {
	fmt.Println("getting total size for " + d.Path + "... ")
	size := 0
	for _, child := range d.ChildDirectories {
		size += child.GetTotalSize()
	}
	for _, file := range d.Files {
		size += file.Size
	}
	fmt.Println("size for " + d.Path + ": " + strconv.Itoa(size))
	return size
}
