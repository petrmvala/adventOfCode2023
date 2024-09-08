package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

type Candidate struct {
	value   int
	x_start int
	x_end   int
}

func sum(schematic string) int {
	var (
		buf        string
		buffering  bool = false
		prevPool   []Candidate
		thisPool   []Candidate
		prevSymbol []int
		thisSymbol []int
		pos_x      int = 0
	)
	fmt.Println(schematic)
	for _, s := range schematic {
		pos_x++
		fmt.Printf(" %q ", s)
		switch {
		case unicode.IsDigit(s):
			buf += string(s)
			buffering = true
		case s == '.':
			if buffering {
				saveInt(&thisPool, pos_x, &buf, &buffering)
			}
		case s == '\n':
			if buffering {
				saveInt(&thisPool, pos_x, &buf, &buffering)
			}
			pos_x = 0
			fmt.Printf("prevPool: %v, thisPool: %v, prevSymbol: %v, thisSymbol: %v\n", prevPool, thisPool, prevSymbol, thisSymbol)
			// check prevPool with thisSymbol
			// check thisPool with thisSymbol
			// check thisPool with prevSymbol
			prevPool = thisPool
			thisPool = []Candidate{}
			prevSymbol = thisSymbol
			thisSymbol = []int{}
		default:
			if buffering {
				saveInt(&thisPool, pos_x, &buf, &buffering)
			}
			thisSymbol = append(thisSymbol, pos_x)
		}
	}
	return 0
}

// func validatePool(pool []Candidate, symbols []int) int {
// 	sum := 0
// 	for _, p := range pool {
// 		// These will never be used as indices so it's ok to get out of bounds
// 		p_min := p.x_start-1
// 		p_max := p.x_end+1
// 		for _, s := range symbols {
// 			if p_min <= s && s <= p_max {

// 			}
// 		}
// 	}
// 	return 0
// }

func saveInt(pool *[]Candidate, x_position int, buffer *string, buffering *bool) {
	tmp, err := strconv.Atoi(*buffer)
	if err != nil {
		log.Fatal(err)
	}
	*pool = append(*pool, Candidate{
		value:   tmp,
		x_start: x_position - len(*buffer),
		x_end:   x_position - 1,
	})
	*buffer = ""
	*buffering = false
}
