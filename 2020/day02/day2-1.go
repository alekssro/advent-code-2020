package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	s "strings"
)

func main() {

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// iterate over lines and count valid passwords
	nValPasswords := 0
	for _, line := range lines {

		// split lines into policy and pswd
		splittedLine := s.Split(line, ":")
		policy := splittedLine[0]
		password := s.TrimSpace(splittedLine[1])

		// parse policy
		splittedPolicy := s.Split(policy, " ")
		numsPolicy := splittedPolicy[0]
		chrPolicy := splittedPolicy[1]
		splittedNumsPolicy := s.Split(numsPolicy, "-")
		pMinReps := toInt(splittedNumsPolicy[0])
		pMaxReps := toInt(splittedNumsPolicy[1])

		// iterate over password to count chr reps
		chrReps := 0
		for _, char := range password {

			if string(char) == chrPolicy {
				chrReps++
			}
		}

		// check if password follows policy (betweem min and max reps)
		// increase nValPasswords if true
		if chrReps >= pMinReps && chrReps <= pMaxReps {
			nValPasswords++
		}
	}
	println(nValPasswords)
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
