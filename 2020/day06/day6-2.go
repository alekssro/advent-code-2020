package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// track total time
	defer timeTrack(time.Now(), "main")

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// init variables
	commonAnswers := ""
	totalYes := 0

	// iterate over lines
	for i, line := range lines {

		// add first line of document to commonAnswers
		if i == 0 {
			commonAnswers = line
			continue
		}

		if line != "" {
			// in group
			// iterate over common answers
			for _, answer := range commonAnswers {
				// if answer not in line, delete answer
				if !runeInString(answer, line) {
					commonAnswers = strings.Replace(commonAnswers,
						string(answer), "", -1)
				}
			}

		} else if line == "" {
			// end of group, sum num common answers and reinit string
			totalYes += len(commonAnswers)
			commonAnswers = lines[i+1]
		}

		if i == len(lines)-1 {
			// sum last group at EOF
			totalYes += len(commonAnswers)
		}

	}
	fmt.Println(totalYes)
}

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Search if rune is in string
func runeInString(a rune, str string) bool {
	for i := range str {
		if a == rune(str[i]) {
			return true
		}
	}
	return false
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
