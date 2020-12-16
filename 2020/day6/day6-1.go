package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	// track total time
	defer timeTrack(time.Now(), "main")

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// iterate over lines
	for i, line := range lines {
		fmt.Println(i, line)

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

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
