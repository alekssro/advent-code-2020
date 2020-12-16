package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	s "strings"
)

func main() {

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// init variables
	mandatoryPassportFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	doc := make(map[string]string)
	nValidIds := 0

	// iterate over lines
	for i, line := range lines {

		if len(line) > 0 {
			// line contains info; save fields to doc map
			fields := s.Split(line, " ")
			for _, field := range fields {
				splittedField := s.Split(field, ":")
				doc[splittedField[0]] = splittedField[1]
			}
		}

		if len(line) == 0 || i == len(lines)-1 {
			// empty line or EOF; end of doc, check if valid
			if isValidDocument(doc, mandatoryPassportFields) {
				nValidIds++
			}
			// Reset doc map
			doc = make(map[string]string)
		}
	}
	fmt.Println(nValidIds)
}

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isInStringArray(a string, arr []string) bool {
	for i := range arr {
		if a == arr[i] {
			return true
		}
	}
	return false
}

func isValidDocument(doc map[string]string, req []string) bool {

	if len(doc) < len(req) {
		// insufficient number of fields in passport
		return false
	}

	// get required keys of document
	keys := make([]string, 0, len(doc))
	for k := range doc {
		if isInStringArray(k, req) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	sort.Strings(req)

	if len(keys) == len(req) {
		for i := range keys {
			if keys[i] != req[i] {
				return false
			}
		}
		return true
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
