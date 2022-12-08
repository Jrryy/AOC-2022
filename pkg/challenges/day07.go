package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
)

type node struct {
	parent *node
	files  []*node
	name   string
	size   int
	isDir  bool
}

var totalFoldersSum = 0

func (n *node) addFolders() int {
	if n.isDir {
		currentSum := 0
		for _, file := range n.files {
			currentSum += file.addFolders()
		}
		if currentSum < 100000 {
			totalFoldersSum += currentSum
		}
		n.size = currentSum
		return currentSum
	} else {
		return n.size
	}
}

func (n *node) findDeletableFolder(minSize int) int {
	if n.isDir {
		maxDeletedSize := 70_000_000
		for _, file := range n.files {
			deletedSize := file.findDeletableFolder(minSize)
			if deletedSize > minSize && deletedSize < maxDeletedSize {
				maxDeletedSize = deletedSize
			}
		}
		if maxDeletedSize != 70_000_000 {
			return maxDeletedSize
		}
		if n.size >= minSize {
			return n.size
		}
	}
	return 0
}

func Day07() {
	input := utils.ImportInputLines(7)

	current := &node{
		parent: nil,
		files:  make([]*node, 0),
		name:   "/",
		size:   0,
		isDir:  true,
	}

	root := current

	line := 1

	for line < len(input) {
		if strings.HasPrefix(input[line], "$ ") {
			if input[line] != "$ ls" {
				folderName := strings.Split(input[line], " ")[2]
				if folderName == ".." {
					current = current.parent
				} else {
					for _, file := range current.files {
						if file.name == folderName {
							current = file
							break
						}
					}
				}
			}
		} else {
			parts := strings.Split(input[line], " ")
			fileSize, notAnInt := strconv.Atoi(parts[0])
			if notAnInt != nil {
				fileSize = 0
			}
			newFile := node{
				parent: current,
				files:  make([]*node, 0),
				name:   parts[1],
				size:   fileSize,
				isDir:  notAnInt != nil,
			}

			current.files = append(current.files, &newFile)
		}
		line++
	}

	totalSpace := root.addFolders()
	spaceToFree := totalSpace - 40_000_000
	deletedSize := root.findDeletableFolder(spaceToFree)
	println(totalFoldersSum)
	println(deletedSize)
}
