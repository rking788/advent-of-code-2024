package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	// 1875
	part1("input.txt")
}

type index struct {
	i int
	j int
}

func part1(filename string) {

	workerCount := 10
	puzzle, err := readPuzzle(filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	indexChan := make(chan index)
	resultChan := make(chan int)

	// Start a goroutine to do the summing of results
	go func(results chan int) {

		total := int64(0)
		for result := range results {
			total += int64(result)
		}

		fmt.Printf("Total = %d\n", total)

	}(resultChan)

	waitGroup := sync.WaitGroup{}

	// Start 10 workers for processing puzzle indexes
	for i := 0; i < workerCount; i++ {

		go func(c chan index, results chan int, p [][]rune, wg *sync.WaitGroup) {
			for {
				coordinates := <-c
				count := checkIndexPart2(coordinates.i, coordinates.j, p)
				results <- count
				wg.Done()
			}
		}(indexChan, resultChan, puzzle, &waitGroup)
	}

	for i := range puzzle {
		for j := range puzzle[i] {
			waitGroup.Add(1)
			indexChan <- index{i, j}
		}
	}

	waitGroup.Wait()
	close(resultChan)

	time.Sleep(1 * time.Millisecond)
	//fmt.Println(puzzle)
}

func boundedIndex(i, j int, puzzle [][]rune) rune {
	if i < 0 || j < 0 {
		return '.'
	} else if i >= len(puzzle) {
		return '.'
	} else if j >= len(puzzle[i]) {
		return '.'
	}
	// 88, 77, 65, 83
	return puzzle[i][j]
}

func checkIndex(i, j int, puzzle [][]rune) int {

	count := 0

	if puzzle[i][j] != 'X' {
		return 0
	}

	// Up
	if boundedIndex(i-1, j, puzzle) == 'M' &&
		boundedIndex(i-2, j, puzzle) == 'A' &&
		boundedIndex(i-3, j, puzzle) == 'S' {
		count++
	}

	// Down
	if boundedIndex(i+1, j, puzzle) == 'M' &&
		boundedIndex(i+2, j, puzzle) == 'A' &&
		boundedIndex(i+3, j, puzzle) == 'S' {
		count++
	}
	// Left
	if boundedIndex(i, j-1, puzzle) == 'M' &&
		boundedIndex(i, j-2, puzzle) == 'A' &&
		boundedIndex(i, j-3, puzzle) == 'S' {
		count++
	}
	// Right
	if boundedIndex(i, j+1, puzzle) == 'M' &&
		boundedIndex(i, j+2, puzzle) == 'A' &&
		boundedIndex(i, j+3, puzzle) == 'S' {
		count++
	}

	// Diagonals
	if boundedIndex(i-1, j-1, puzzle) == 'M' &&
		boundedIndex(i-2, j-2, puzzle) == 'A' &&
		boundedIndex(i-3, j-3, puzzle) == 'S' {
		count++
	}
	if boundedIndex(i-1, j+1, puzzle) == 'M' &&
		boundedIndex(i-2, j+2, puzzle) == 'A' &&
		boundedIndex(i-3, j+3, puzzle) == 'S' {
		count++
	}
	if boundedIndex(i+1, j+1, puzzle) == 'M' &&
		boundedIndex(i+2, j+2, puzzle) == 'A' &&
		boundedIndex(i+3, j+3, puzzle) == 'S' {
		count++
	}
	if boundedIndex(i+1, j-1, puzzle) == 'M' &&
		boundedIndex(i+2, j-2, puzzle) == 'A' &&
		boundedIndex(i+3, j-3, puzzle) == 'S' {
		count++
	}

	return count
}

func checkIndexPart2(i, j int, puzzle [][]rune) int {

	count := 0

	if puzzle[i][j] != 'A' {
		return 0
	}

	// Diagonals
	// M M
	//  A
	// S S
	if boundedIndex(i-1, j-1, puzzle) == 'M' &&
		boundedIndex(i-1, j+1, puzzle) == 'M' &&
		boundedIndex(i+1, j-1, puzzle) == 'S' &&
		boundedIndex(i+1, j+1, puzzle) == 'S' {
		count++
	}
	// S M
	//  A
	// S M
	if boundedIndex(i-1, j+1, puzzle) == 'M' &&
		boundedIndex(i+1, j+1, puzzle) == 'M' &&
		boundedIndex(i-1, j-1, puzzle) == 'S' &&
		boundedIndex(i+1, j-1, puzzle) == 'S' {
		count++
	}
	// M S
	//  A
	// M S
	if boundedIndex(i-1, j-1, puzzle) == 'M' &&
		boundedIndex(i+1, j-1, puzzle) == 'M' &&
		boundedIndex(i-1, j+1, puzzle) == 'S' &&
		boundedIndex(i+1, j+1, puzzle) == 'S' {
		count++
	}
	// S S
	//  A
	// M M
	if boundedIndex(i+1, j-1, puzzle) == 'M' &&
		boundedIndex(i+1, j+1, puzzle) == 'M' &&
		boundedIndex(i-1, j-1, puzzle) == 'S' &&
		boundedIndex(i-1, j+1, puzzle) == 'S' {
		count++
	}

	// Apparently these cases are invalid, I guess MAS needs to be spelled in the same direction for the pair.
	// Adding them in leads to an answer that is too high.

	// S M
	// 	A
	// M S
	//if boundedIndex(i-1, j+1, puzzle) == 'M' &&
	//	boundedIndex(i+1, j-1, puzzle) == 'M' &&
	//	boundedIndex(i-1, j-1, puzzle) == 'S' &&
	//	boundedIndex(i+1, j+1, puzzle) == 'S' {
	//	count++
	//}
	// M S
	//  A
	// S M
	//if boundedIndex(i-1, j-1, puzzle) == 'M' &&
	//	boundedIndex(i+1, j+1, puzzle) == 'M' &&
	//	boundedIndex(i-1, j+1, puzzle) == 'S' &&
	//	boundedIndex(i+1, j-1, puzzle) == 'S' {
	//	count++
	//}

	return count
}

func readPuzzle(filename string) ([][]rune, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := make([][]rune, 0)
	for {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		lineRunes := make([]rune, 0, len(line))
		for _, r := range line {
			lineRunes = append(lineRunes, r)
		}

		result = append(result, lineRunes)
	}

	return result, nil
}
