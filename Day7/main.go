package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	parent *Directory
	name   string
	size   int64
}

type Directory struct {
	parent              *Directory
	name                string
	internalDirectories map[string]*Directory
	files               map[string]*File
	size                int64
}

func (wd *Directory) calculateDirectorySize() {
	for _, dir := range wd.internalDirectories {
		dir.calculateDirectorySize()
	}
	for _, file := range wd.files {
		wd.size = wd.size + file.size
	}
	for _, dir := range wd.internalDirectories {
		wd.size = wd.size + dir.size
	}
}
func (wd *Directory) collectDirectoriesWithSize(minSize, maxSize int64) []*Directory {
	var directories []*Directory

	for _, dir := range wd.internalDirectories {
		directories = append(directories, dir.collectDirectoriesWithSize(minSize, maxSize)...)
	}

	if wd.size >= minSize && wd.size <= maxSize {
		directories = append(directories, wd)
	}

	return directories
}

func main() {
	directory := loadFilesData()
	directory = executeCommand("$ cd /", directory)
	directory.calculateDirectorySize()

	answer1(directory)
}

func answer1(workingDirectory *Directory) {
	dirWithSize10000 := workingDirectory.collectDirectoriesWithSize(0, 100000)
	var size int64

	for _, dir := range dirWithSize10000 {
		size += dir.size
	}

	fmt.Printf("Sum of size of all directories less than 10000 is %d", size)
}

func loadFilesData() *Directory {
	workingDirectory := &Directory{
		parent:              nil,
		name:                "/",
		internalDirectories: make(map[string]*Directory),
		files:               make(map[string]*File),
		size:                0,
	}

	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if isCommand(line) {
			workingDirectory = executeCommand(line, workingDirectory)
		} else {
			loadData(line, workingDirectory)
		}
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return workingDirectory
}

func isCommand(input string) bool {
	return len(input) > 0 && input[0] == '$'
}
func isDirectory(input string) bool {
	return len(input) > 0 && input[0:3] == "dir"
}

func executeCommand(input string, workingDirectory *Directory) *Directory {
	command := input[2:4]
	if command == "cd" {
		path := input[5:]
		workingDirectory = changeDirectory(workingDirectory, path)
	}
	return workingDirectory
}

func changeDirectory(workingDirectory *Directory, path string) *Directory {
	if path == ".." {
		workingDirectory = workingDirectory.parent
	} else if path == "/" {
		for {
			if workingDirectory.parent == nil || workingDirectory.name == "/" {
				break
			}

			workingDirectory = workingDirectory.parent
		}
	} else {
		dir := workingDirectory.internalDirectories[path]
		if dir != nil {
			workingDirectory = dir
		} else {
			println("Directory not changed as it does not exists")
		}
	}

	return workingDirectory
}

func loadData(input string, workingDirectory *Directory) {
	if isDirectory(input) {
		addDirectory(input, workingDirectory)
	} else {
		addFile(input, workingDirectory)
	}
}

func addFile(input string, workingDirectory *Directory) {
	split := strings.Fields(input)
	size, _ := strconv.Atoi(split[0])
	name := split[1]
	file := &File{
		parent: workingDirectory,
		name:   name,
		size:   int64(size),
	}

	workingDirectory.files[name] = file
}

func addDirectory(input string, workingDirectory *Directory) {
	name := input[4:]
	childDirectory := &Directory{
		name:                name,
		parent:              workingDirectory,
		internalDirectories: make(map[string]*Directory),
		files:               make(map[string]*File),
		size:                0,
	}

	workingDirectory.internalDirectories[name] = childDirectory
}
