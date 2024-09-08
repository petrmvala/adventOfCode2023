package main

import "testing"

var schematic = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestSum(t *testing.T) {
	got := sum(schematic)
	want := 4361
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
