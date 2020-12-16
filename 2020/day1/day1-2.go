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
	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j {
				break
			}
			for k, line3 := range lines {
				// skip at same line
				if i == k || j == k {
					break
				}

				// transform to int
				num1 := toInt(line1)
				num2 := toInt(line2)
				num3 := toInt(line3)

				// check if sum equal 2020
				if num1+num2+num3 == 2020 {
					// fmt.Println(num1, num2, num3)
					// print result of multiplying (submission)
					fmt.Println(num1 * num2 * num3)
				}
			}
		}
	}
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
