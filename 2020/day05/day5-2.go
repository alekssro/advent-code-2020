package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	nRows    = 128 // number of rows
	nCols    = 8   // number of seats per row
	nChrsRow = 7   // number of chr for row instructions
)

func main() {

	// track total time
	defer timeTrack(time.Now(), "main")

	// First: Read lines of the input files into an array
	lines, err := readLines("input.txt")
	check(err)

	// init vars
	var (
		toRow   = ""
		toCol   = ""
		row     = -1
		col     = -1
		seatID  = 0
		seatIDs = []int{}
	)

	// iterate over lines and save seatIDs
	for _, line := range lines {

		// Split Get-to-Row indications and Get-to-Seat indications
		toRow, toCol = line[:nChrsRow], line[nChrsRow:]
		row = binSearch(toRow, nRows, 'F', 'B')
		col = binSearch(toCol, nCols, 'L', 'R')

		// calculate seatID
		seatID = row*8 + col

		// // debug
		// fmt.Println(row, col, seatID)
		seatIDs = append(seatIDs, seatID)

	}
	sort.Ints(seatIDs)

	// find missing seatID (mine)
	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] != seatIDs[i]+1 {
			fmt.Println(seatIDs[i] + 1)
		}
	}

	// time.Sleep(100 * time.Millisecond)
}

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func binSearch(instr string, n int, takeLow, takeUp rune) int {

	interval := [2]float64{0., float64(n - 1)}

	for _, dir := range instr {

		if dir == takeLow {
			// we keep first half of interval
			interval[1] -= math.Floor((interval[1] - interval[0]) / 2)
		} else if dir == takeUp {
			// we keep last half of interval
			interval[0] += math.Ceil((interval[1] - interval[0]) / 2)
		} else {
			raiseInvInstruction(dir, []rune{takeUp, takeLow})
		}
	}
	// return mid point
	return int((interval[1] + interval[0]) / 2)
}

func raiseInvInstruction(ins rune, validIns []rune) {
	// unknown instruction
	s := fmt.Sprintf("Unknown instruction found: '%s' \nAllowed instructions: ['%s', '%s']",
		string(ins), string(validIns[0]), string(validIns[1]))
	panic(s)
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
