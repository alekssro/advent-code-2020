package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	s "strings"
	"time"
)

func main() {

	// track total time
	defer timeTrack(time.Now(), "main")

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

func checkValidValues(doc map[string]string) bool {

	checks := []bool{}
	checks = append(checks,
		checkYearInRange(doc["byr"], 1920, 2002),
		checkYearInRange(doc["iyr"], 2010, 2020),
		checkYearInRange(doc["eyr"], 2020, 2030),
		checkHeight(doc["hgt"]),
		checkHairColor(doc["hcl"]),
		checkEyeColor(doc["ecl"]),
		checkPID(doc["pid"]),
	)

	// Pass all checks
	for _, check := range checks {
		if !check {
			// found one failed test
			return false
		}
	}
	return true
}

func checkEyeColor(ecl string) bool {

	validColors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, vc := range validColors {
		if ecl == vc {
			// eye color is a valid color
			return true
		}
	}

	return false
}

func checkPID(pid string) bool {

	if len(pid) != 9 {
		// invalid length of pid
		return false
	}

	return isInt(pid)

}

func checkHairColor(hcl string) bool {

	if hcl[:1] != "#" {
		// does not start with "#"; invalid
		return false
	}

	if len(hcl[1:]) != 6 {
		// "#" not followed by exactly six characters; invalid
		return false
	}

	// check every character to be 0-9 or a-f
	for _, chr := range hcl[1:] {
		if !((chr >= '0' && chr <= '9') || (chr >= 'a' && chr <= 'f')) {
			// if not between 0-9 or a-f; invalid
			return false
		}
	}

	return true
}

func checkHeight(h string) bool {

	units := h[len(h)-2:]

	switch units {
	case "cm":
		// height in centimeters
		value := toInt(h[:len(h)-2])
		if value >= 150 && value <= 193 {
			return true
		}
	case "in":
		// height in inches
		value := toInt(h[:len(h)-2])
		if value >= 59 && value <= 76 {
			return true
		}
	}

	return false // if not in cases

}

func checkYearInRange(y string, startY, endY int) bool {

	year := toInt(y)

	if year >= startY && year <= endY {
		return true
	}

	return false

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
		return checkValidValues(doc)
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

func isInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func toInt(s string) int {
	// transform string to int
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Warning: not able to convert string to int\n", err)
	}

	return num
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
