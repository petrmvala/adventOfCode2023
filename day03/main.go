package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

type Candidate struct {
	value   int
	x_start int
	x_end   int
}

func main() {

	readFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	s := bufio.NewScanner(readFile)
	s.Split(bufio.ScanLines)

	var schema string

	for s.Scan() {
		schema += fmt.Sprintln(s.Text()) // Println will add back the final '\n'
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	start := time.Now()
	sum := sum(schema)
	timeElapsed := time.Since(start)
	fmt.Println("result: ", sum, "\ntime elapsed: ", timeElapsed)

}

func sum(schematic string) int {
	var (
		valid               bool
		pos_x, sum          int
		buffer              string
		abovePool, thisPool []Candidate
	)
	thisSymbol := make(map[int]bool)
	aboveSymbol := thisSymbol

	pos_x = -1
	for _, s := range schematic {
		pos_x++
		switch {
		case unicode.IsDigit(s):
			buffer += string(s)
			if aboveSymbol[pos_x] {
				valid = true
			}
		case s == '.':
			if len(buffer) > 0 {
				if aboveSymbol[pos_x] {
					valid = true
				}
				c := getCandidate(buffer, pos_x-1)
				buffer = ""
				if valid {
					sum += c.value
				} else {
					thisPool = append(thisPool, c)
				}
			}
			valid = aboveSymbol[pos_x]
		case s == '\n':
			if len(buffer) > 0 {
				c := getCandidate(buffer, pos_x-1)
				buffer = ""
				if valid {
					sum += c.value
				} else {
					thisPool = append(thisPool, c)
				}
			}
			abovePool, thisPool = thisPool, []Candidate{}
			aboveSymbol, thisSymbol = thisSymbol, map[int]bool{}
			valid = false
			pos_x = -1
		default:
			thisSymbol[pos_x] = true
			valid = true
			if len(buffer) > 0 {
				c := getCandidate(buffer, pos_x-1)
				buffer = ""
				sum += c.value
			}
			for _, c := range abovePool {
				if c.x_start-1 > pos_x {
					continue
				}
				if c.x_end+1 < pos_x {
					continue
				}
				sum += c.value
			}
		}
	}

	return sum
}

func getCandidate(s string, pos_x int) Candidate {
	tmp, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return Candidate{
		value:   tmp,
		x_start: pos_x - (len(s) - 1),
		x_end:   pos_x,
	}
}
