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
		pos1 := toInt(splittedNumsPolicy[0]) - 1 // index 1 to zero
		pos2 := toInt(splittedNumsPolicy[1]) - 1 // index 1 to zero

		// check if pos1 or pos2 contains character (not in both)
		if password[pos1] != password[pos2] {
			if string(password[pos1]) == chrPolicy || string(password[pos2]) == chrPolicy {
				nValPasswords++
			}
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
