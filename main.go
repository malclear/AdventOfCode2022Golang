package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/cast"
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
	//inFileAddMode := false

	dirList := make(map[string]*Directory)
	pwd := NewDirectoryByName("/", nil)
	dirList[pwd.Path] = pwd

	for _, line := range parsed {
		if strings.HasPrefix(line, "$ cd ..") {
			pathArray := strings.Split(pwd.Path, "/")
			pathArray = pathArray[1:]                     // removes first, empty string
			pathArray = pathArray[:len(pathArray)-1]      // removes the topmost directory
			newPath := "/" + strings.Join(pathArray, "/") // rejoins everything with a leading slash
			pwd = getDirectory(newPath, dirList)
		} else if strings.HasPrefix(line, "$ cd ") {
			// This handles only relative paths
			newDir := strings.Replace(line, "$ cd ", "", -1)
			newDir = strings.Replace(newDir, "$ cd ", "/", -1)
			if newDir == "/" {
				pwd = dirList["/"]
				continue
			}
			path := pwd.Path
			if path == "/" {
				path = ""
			}
			pwd = dirList[path+"/"+newDir]

		} else if strings.HasPrefix(line, "$ ls") {
			// Nothing to do here!!
			//inFileAddMode = true
		} else if strings.HasPrefix(line, "dir ") {
			dirName := strings.Replace(line, "dir ", "", -1)
			newDir := NewDirectoryByName(dirName, pwd)
			pwd.ChildDirectories = append(pwd.ChildDirectories, newDir)
			dirList[newDir.Path] = newDir
		} else {
			fileListing := strings.Split(line, " ")
			addNewFile(fileListing[1], cast.ToInt(fileListing[0]), pwd)
		}
	}

	smallDirTotalSizes := 0
	for _, directory := range dirList {

		dirSize := directory.GetTotalSize()
		if dirSize <= 100000 {
			smallDirTotalSizes += dirSize
		}
	}

	return smallDirTotalSizes
}

func getDirectory(pathToDirectory string, dirList map[string]*Directory) *Directory {
	return dirList[pathToDirectory]
}

func addNewFile(name string, size int, pwd *Directory) {
	var newFile File
	newFile.Init(name, size)
	pwd.Files = append(pwd.Files, &newFile)

}

//func changeDirectory(path string, pwd *string) string {
//	*pwd = *pwd + path
//	return *pwd + path
//}

func part2(input string) int {
	parsed := parseInput(input)
	//inFileAddMode := false

	dirList := make(map[string]*Directory)
	pwd := NewDirectoryByName("/", nil)
	dirList[pwd.Path] = pwd

	for _, line := range parsed {
		if strings.HasPrefix(line, "$ cd ..") {
			pathArray := strings.Split(pwd.Path, "/")
			pathArray = pathArray[1:]                     // removes first, empty string
			pathArray = pathArray[:len(pathArray)-1]      // removes the topmost directory
			newPath := "/" + strings.Join(pathArray, "/") // rejoins everything with a leading slash
			pwd = getDirectory(newPath, dirList)
		} else if strings.HasPrefix(line, "$ cd ") {
			// This handles only relative paths
			newDir := strings.Replace(line, "$ cd ", "", -1)
			newDir = strings.Replace(newDir, "$ cd ", "/", -1)
			if newDir == "/" {
				pwd = dirList["/"]
				continue
			}
			path := pwd.Path
			if path == "/" {
				path = ""
			}
			pwd = dirList[path+"/"+newDir]

		} else if strings.HasPrefix(line, "$ ls") {
			// Nothing to do here!!
			//inFileAddMode = true
		} else if strings.HasPrefix(line, "dir ") {
			dirName := strings.Replace(line, "dir ", "", -1)
			newDir := NewDirectoryByName(dirName, pwd)
			pwd.ChildDirectories = append(pwd.ChildDirectories, newDir)
			dirList[newDir.Path] = newDir
		} else {
			fileListing := strings.Split(line, " ")
			addNewFile(fileListing[1], cast.ToInt(fileListing[0]), pwd)
		}
	}

	spaceRequired := 30000000
	diskSize := 70000000
	spaceUsed := dirList["/"].GetTotalSize()
	amountToFree := spaceRequired - (diskSize - spaceUsed)

	//pathToSmallest := "/"
	smallestSize := spaceUsed
	for _, directory := range dirList {
		dirSize := directory.GetTotalSize()
		if dirSize < smallestSize && dirSize > amountToFree {
			smallestSize = dirSize
			//pathToSmallest = directory.Path
		}
	}
	return smallestSize
}
