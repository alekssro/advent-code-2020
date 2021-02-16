package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type content struct {
	bag string
	num int
}
type rule struct {
	bag     string
	content []content
}

var rules map[string]rule

func main() {

	// track total time
	defer timeTrack(time.Now(), "main")

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// init container bags
	rules = make(map[string]rule)
	result := 0

	// Save bag rules
	for _, line := range lines {

		// match container bag - contained bags separation
		rule := toRule(line)
		rules[rule.bag] = rule

	}

	// Count bags that "shiny gold" contains
	result = countNestedBags("shiny gold", rules) - 1
	fmt.Println(result)

}

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toRule(s string) (r rule) {
	re := regexp.MustCompile(`^(?P<Bag>.+) bags contain (?P<Content>.+)\.$`)
	if !re.MatchString(s) {
		panic("Did not recognize " + s)
	}
	res := re.FindStringSubmatch(s)
	r.bag = res[1] // first submatch
	r.content = make([]content, 0)
	if res[2] == "no other bags" {
		return r
	}
	cre := regexp.MustCompile(`^(?P<Count>\d+) (?P<Bag>.+) (bags|bag)$`)
	contents := strings.Split(res[2], ", ")
	for _, c := range contents {
		contentFields := cre.FindStringSubmatch(c)
		r.content = append(r.content, content{
			bag: contentFields[2],
			num: toInt(contentFields[1]),
		})
	}
	return r

}

// Recursively count the number of bags inside bags
func countNestedBags(bag string, rules map[string]rule) int {

	count := 1
	rule := rules[bag]

	for _, c := range rule.content {
		count += c.num * countNestedBags(c.bag, rules)
	}

	return count
}

// Recursively determine if can carry "shiny gold" bag
func canCarryGoldBag(bag string) bool {
	if bag == "shiny gold" {
		return true
	}
	rule, hasRule := rules[bag]
	if !hasRule {
		panic("Could not find rule for " + bag)
	}
	if len(rule.content) == 0 {
		return false
	}
	for _, c := range rule.content {
		if canCarryGoldBag(c.bag) {
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
