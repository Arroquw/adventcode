package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type fileSize int64
type fileSizes []fileSize

type Dir struct {
	Name		string
	ParentDir	*Dir
	ChildDirs	map[string]*Dir
	Files		map[string]fileSize
	TotalSize	fileSize
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	day7(input)
}

func findCandidates(itr *Dir) fileSize {
	SizeLimit := fileSize(100000)

	sum := fileSize(0)
	if itr.TotalSize <= SizeLimit {
		sum += itr.TotalSize
	}
	for _, child := range itr.ChildDirs {
		sum += findCandidates(child)
	}
	return sum
}

func (candidates *fileSizes) daypt2(d *Dir, s fileSize) {
	SpaceRequired := fileSize(30000000) - s
	for _, val := range d.ChildDirs {
		if val.TotalSize >= SpaceRequired {
			*candidates = append(*candidates, val.TotalSize)
		}
		candidates.daypt2(val, s)
	}
}

func day7(input []string) {
	root := &Dir{
		Name: "/",
		ChildDirs: map[string]*Dir{},
	}
	iter := root

	for _, c := range input {
		switch c[0:1] {
		case "$":
			if c == "$ ls" {
				continue
			} else {
				temp := strings.Split(c, " ")[2]
				if temp == ".." {
					iter = iter.ParentDir
				} else {
					if _, ok := iter.ChildDirs[temp]; !ok {
						iter.ChildDirs[temp] = &Dir {
							Name: temp,
							ParentDir: iter,
							ChildDirs: map[string]*Dir{},
							Files: map[string]fileSize{},
						}
					}
					iter = iter.ChildDirs[temp]
				}
			}
		default:
			if strings.HasPrefix(c, "dir") {
				temp := strings.Split(c, " ")[1]
				if _, ok := iter.ChildDirs[temp]; !ok {
					iter.ChildDirs[temp] = &Dir {
						Name: temp,
						ParentDir: iter,
						ChildDirs: map[string]*Dir{},
						Files: map[string]fileSize{},
					}
				}
			} else {
				temp := strings.Split(c, " ")
				s, _ := strconv.ParseInt(temp[0], 10, 64)
				iter.Files[temp[1]] = fileSize(s)
			}

		}
	}
	addSizes(root)
	fmt.Println(findCandidates(root))
	CurrentSpace := fileSize(70000000 - root.TotalSize)
	s := make(fileSizes, 0)
	s.daypt2(root, CurrentSpace)
	for _, c := range s {
		if s[1] > c {
			s[0] = c
			s[1] = s[0]
		} else {
			s[0] = s[1]
		}
	}
	fmt.Println(s[0])
}

/* 
func day1(scanner *bufio.Scanner) {
        var temp2 []int64
        var temp string
        max := (int64)(0)
        sum := make([]int64, 2)
        scanValid := true
        for {
                if scanValid {
                        scanValid = scanner.Scan()
                        temp = scanner.Text()
                }
                if temp != "" {
                        num, _ := strconv.ParseInt(temp, 10, 64)
                        temp2 = append(temp2, num)
                        continue
                }
                for _, ijk := range temp2 {
                        sum[0] += ijk
                }
                if sum[1] < sum[0] {
                        max = sum[0]
                        sum[1] = sum[0]
                } else {
                        max = sum[1]
                }
                sum[0] = 0
                temp2 = make([]int64, 0)
                if !scanValid {
                        break
                }
        }

        fmt.Println(max)
        if scanner.Err() != nil {
                fmt.Println("error")
        }
}
*/

func addSizes(itr *Dir) fileSize {
	totalSize := fileSize(0)

	for _, childItr := range itr.ChildDirs {
		totalSize += addSizes(childItr)
	}

	for _, sz := range itr.Files {
		totalSize += sz
	}

	itr.TotalSize = totalSize

	return totalSize

}
