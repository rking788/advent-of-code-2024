package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	mulPattern = regexp.MustCompile(`mul\((?P<first>\d{1,3}),(?P<second>\d{1,3})\)|do\(\)|don't\(\)`)
)

func main() {
	//part1("input.txt")

	// 113,908,771
	// 111,762,583
	part2("input.txt")
}

func part1(filename string) {

	lines, err := readInput(filename)
	if err != nil {
		log.Fatalf("failed to read input [err:%s]\n", err.Error())
	}

	sum := int64(0)
	for _, l := range lines {

		lineTotal, _, err := processLine(l, true)
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += lineTotal
	}

	fmt.Printf("Answer = %d\n", sum)
}

func part2(filename string) {

	lines, err := readInput(filename)
	if err != nil {
		log.Fatalf("failed to read input [err:%s]\n", err.Error())
	}

	sum := int64(0)
	operationsEnabled := true
	for _, l := range lines {

		var lineTotal int64
		lineTotal, operationsEnabled, err = processLine(l, operationsEnabled)
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += lineTotal
	}

	fmt.Printf("Answer = %d\n", sum)
}

func processLine(line string, operationsEnabled bool) (int64, bool, error) {

	sum := int64(0)
	matches := mulPattern.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		//fmt.Printf("Matches -- %+v\n", match)
		// Index 0 is the complete regex match not the named groups

		if strings.HasPrefix(match[0], "mul") && operationsEnabled {
			first, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, false, err
			}
			second, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, false, err
			}

			sum += int64(first * second)
		} else if match[0] == "do()" {
			operationsEnabled = true
		} else if match[0] == "don't()" {
			operationsEnabled = false
		} else {
			fmt.Printf("Unrecognized operation -- %s\n", match[0])
		}
	}

	return sum, operationsEnabled, nil
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
