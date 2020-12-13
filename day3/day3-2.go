package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// Count number of trees with different slopes
	nTrees := countTrees(lines, 1, 1) * countTrees(lines, 3, 1) *
		countTrees(lines, 5, 1) * countTrees(lines, 7, 1) * countTrees(lines, 1, 2)
	fmt.Println(nTrees)
}

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countTrees(lines []string, right int, down int) int {

	nLines := len(lines)
	lenLine := len(lines[0])
	nTrees := 0
	pos := [2]int{0, 0}
	for i := 0; i < nLines-down; i += down {
		pos[0] = (pos[0] + right) % lenLine // reset if end of line
		pos[1] = pos[1] + down
		if lines[pos[1]][pos[0]] == '#' {
			nTrees++
		}
	}
	return nTrees
}

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return lines, err
}

func toInt(s string) int {
	// transform to int
	num, err := strconv.Atoi(s)
	check(err)

	return num
}
