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

// Instruction represents a machine instruction: a command and a value
type Instruction struct {
	Cmd  string
	Val  int
	Done bool
}

// Machine is a list of instructions that can be run
type Machine []Instruction

// RunEndBeforeRepeat implements running the instructions and either stop before
// repeating itself, or when the instruction after last will be run (program terminated)
func (m Machine) RunEndBeforeRepeat() (int, bool, error) {
	var acc, i int

	// exec instructions until current intruction was visited or ended program
	for {
		if i == len(m) {
			return acc, true, nil
		}
		if i < 0 || i >= len(m) {
			err := fmt.Errorf("index %d outside of machine length (%d)", i, len(m))
			return 0, false, err
		}
		if m[i].Done {
			return acc, false, nil
		}

		// set current instruction to visited
		m[i].Done = true

		// exec corresponding instruction
		instr := m[i]
		switch instr.Cmd {

		case "nop":
			i++

		case "acc":
			acc += instr.Val
			i++

		case "jmp":
			i += instr.Val

		default:
			panic("Instruction Error: Unrecognize instruction.")
		}
	}
}

func main() {

	// track total time
	defer timeTrack(time.Now(), "main")

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// iterate over lines to save instructions
	instructions, err := getInstructions(lines)
	check(err)

	// init Machine
	m := Machine(instructions)
	acc, finished, err := m.RunEndBeforeRepeat()
	check(err)
	if finished {
		log.Fatal("Finished without repeating")
	}

	fmt.Println(acc)

}

// Saves instructions slice
func getInstructions(lines []string) (Machine, error) {

	var m Machine
	for _, line := range lines {

		splits := strings.Split(line, " ")
		if len(splits) != 2 {

			return m, fmt.Errorf("can't split %s", line)
		}

		cmd := strings.TrimSpace(splits[0])
		val := toInt(splits[1])

		m = append(m, Instruction{Cmd: cmd, Val: val})
	}
	return m, nil
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
