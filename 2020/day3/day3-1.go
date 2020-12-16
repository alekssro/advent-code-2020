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

	// iterate over lines
	nLines := len(lines)
	lenLine := len(lines[0])
	nTrees := 0
	slope := [2]int{3, 1}
	pos := [2]int{0, 0}
	for i := 0; i < nLines-1; i++ {
		pos[0] = (pos[0] + slope[0]) % lenLine // reset if end of line
		pos[1] = pos[1] + slope[1]
		if lines[pos[1]][pos[0]] == '#' {
			nTrees++
		}
	}
	fmt.Println(nTrees)
}

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
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
