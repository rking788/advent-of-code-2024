package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	//part1()
	part2()
}

func part1() {

	//left, right, err := readInput("sample.txt")
	left, right, err := readInput("input.txt")
	if err != nil {
		os.Exit(1)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := int64(0)
	for i := range left {
		distance := math.Abs(float64(left[i] - right[i]))
		sum += int64(distance)
	}

	fmt.Printf("%d\n", sum)
}

func part2() {
	left, counts, err := readInputPart2("input.txt")
	if err != nil {
		os.Exit(1)
	}

	total := int64(0)
	for _, leftValue := range left {
		if count, ok := counts[leftValue]; ok {
			total += int64(count * leftValue)
		}
	}

	fmt.Printf("%d\n", total)
}

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	left, right := make([]int, 0), make([]int, 0)
	isLeft := true
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for {
		if !scanner.Scan() {
			break
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, err
		}

		if isLeft {
			left = append(left, val)
		} else {
			right = append(right, val)
		}

		isLeft = !isLeft
	}

	return left, right, nil
}

func readInputPart2(filename string) ([]int, map[int]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	left := make([]int, 0)
	counts := make(map[int]int)
	isLeft := true
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for {
		if !scanner.Scan() {
			break
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, err
		}

		if isLeft {
			left = append(left, val)
		} else {
			if existing, ok := counts[val]; ok {
				counts[val] = existing + 1
			} else {
				counts[val] = 1
			}
		}

		isLeft = !isLeft
	}

	return left, counts, nil
}
