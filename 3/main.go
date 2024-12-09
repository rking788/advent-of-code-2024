package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	mulPattern = regexp.MustCompile(`mul\((?P<first>\d{1,3}),(?P<second>\d{1,3})\)`)
)

func main() {
	part1("input.txt")
}

func part1(filename string) {

	lines, err := readInput(filename)
	if err != nil {
		log.Fatalf("failed to read input [err:%s]\n", err.Error())
	}

	sum := int64(0)
	for _, l := range lines {

		lineTotal, err := processLine(l)
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += lineTotal
	}

	fmt.Printf("Answer = %d\n", sum)
}

func processLine(line string) (int64, error) {

	sum := int64(0)
	matches := mulPattern.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		//fmt.Printf("Matches -- %+v\n", match)
		// Index 0 is the complete regex match not the named groups
		first, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		second, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		sum += int64(first * second)
	}

	return sum, nil
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := make([]string, 0)
	for {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		result = append(result, line)
	}

	return result, nil
}
