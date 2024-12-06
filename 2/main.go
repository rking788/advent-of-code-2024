package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	maxDifference = 3
)

func main() {
	part1()
	//part2()
}

func part1() {

	values, err := readInput("input.txt")
	if err != nil {
		os.Exit(1)
	}

	count := 0
	for _, line := range values {

		if isLineSafe(line, true) {
			fmt.Printf("safe\n")
			count++
		} else {
			fmt.Printf("not safe\n")
		}

	}

	fmt.Println(count)
}

func part2() {

	// Ended up just using the same isLineSafe method for both parts
}

func readInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := make([][]int, 0)
	for {
		if !scanner.Scan() {
			break
		}

		lineValues := make([]int, 0)
		line := scanner.Text()
		pieces := strings.Split(line, " ")
		for _, piece := range pieces {
			val, err := strconv.Atoi(piece)
			if err != nil {
				return nil, err
			}

			lineValues = append(lineValues, val)
		}

		result = append(result, lineValues)
	}

	return result, nil
}

func isLineSafe(line []int, allowViolation bool) bool {

	fmt.Printf("Checking if line is safe: %+v...", line)
	if len(line) < 2 {
		return true
	}

	isIncreasing := true
	if line[1] < line[0] {
		isIncreasing = false
	}

	violations := make(map[int]int)
	for i := 1; i < len(line); i++ {

		invalid := false
		difference := line[i] - line[i-1]
		if difference == 0 || math.Abs(float64(difference)) > maxDifference {
			invalid = true
		} else if isIncreasing && difference < 0 {
			invalid = true
		} else if !isIncreasing && difference > 0 {
			invalid = true
		}

		if invalid {
			if leftCount, ok := violations[i-1]; ok {
				violations[i-1] = leftCount + 1
			} else {
				violations[i-1] = 1
			}
			if rightCount, ok := violations[i]; ok {
				violations[i] = rightCount + 1
			} else {
				violations[i] = 1
			}
		}
	}

	if len(violations) == 0 {
		return true
	}

	if allowViolation && len(violations) > 0 {

		waitGroup := sync.WaitGroup{}
		safeChan := make(chan bool, len(line))
		for i := range line {

			waitGroup.Add(1)
			go func(wg *sync.WaitGroup, i int, c chan<- bool) {
				defer wg.Done()

				adjusted := make([]int, 0, len(line)-1)
				adjusted = append(adjusted, line[:i]...)
				adjusted = append(adjusted, line[i+1:]...)

				isNowSafe := isLineSafe(adjusted, false)
				c <- isNowSafe
			}(&waitGroup, i, safeChan)
		}

		waitGroup.Wait()
		close(safeChan)

		for isSafe := range safeChan {
			if isSafe {
				return true
			}
		}
	}

	return false
}
