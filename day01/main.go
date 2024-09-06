package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("day01/data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	s := bufio.NewScanner(readFile)
	s.Split(bufio.ScanLines)

	sum := 0
	for s.Scan() {
		sum += lineMatch(s.Text())
	}

	fmt.Println("result: ", sum)
}

func lineMatch(line string) int {
	re, err := regexp.Compile("^\\D*(\\d)(.*(\\d))?\\D*$")
	if err != nil {
		log.Fatalln("failed to compile: ", err)
	}

	match := re.FindStringSubmatch(line)
	// fmt.Println(match)
	if len(match) != 4 {
		log.Fatalln("inconsistent matches: input: ", line, " match: ", match)
	}

	s := match[1] + match[1]
	if len(match[3]) != 0 {
		s = match[1] + match[3]
	}

	res, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("failed to convert: ", err)
	}

	return res
}
