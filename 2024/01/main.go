package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l := newListsData()
	err := l.parseInputFile("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %v\n", err)
	}

	l.calculateTotalDistance()
	fmt.Printf("Total distance between lists: %d\n", l.TotalDistance)

	l.calculateSimilarityScore()
	fmt.Printf("Similarity Score: %d\n", l.SimilarityScore)
}

type ListsData struct {
	List1           []int
	List2           []int
	TotalDistance   int
	List2Occurences map[int]int
	SimilarityScore int
}

func newListsData() *ListsData {
	return &ListsData{
		List1:           make([]int, 0),
		List2:           make([]int, 0),
		TotalDistance:   0,
		List2Occurences: make(map[int]int),
		SimilarityScore: 0,
	}
}

func parseLine(line string) (int, int, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected 2 numbers, got %d", len(parts))
	}

	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}

func (l *ListsData) parseInputFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num1, num2, err := parseLine(line)
		if err != nil {
			return err
		}

		l.List1 = append(l.List1, num1)
		l.List2 = append(l.List2, num2)
	}

	if len(l.List1) != len(l.List2) {
		return fmt.Errorf("list1 length (%d) does not match list2 length (%d)", len(l.List1), len(l.List2))
	}

	return nil
}

func (l *ListsData) calculateTotalDistance() {
	sort.Ints(l.List1)
	sort.Ints(l.List2)

	for i := range len(l.List1) {
		a := l.List1[i]
		b := l.List2[i]

		var diff int
		if a > b {
			diff = a - b
		} else {
			diff = b - a
		}

		l.TotalDistance += diff
	}
}

func (l *ListsData) calculateSimilarityScore() {
	for _, num := range l.List2 {
		l.List2Occurences[num]++
	}

	for _, num := range l.List1 {
		occurences := l.List2Occurences[num]
		l.SimilarityScore += num * occurences
	}
}
