package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left, right []int
	var totalDistance, similarityScore int
	rightOccurences := make(map[int]int)

	parseLine := func(line string) (int, int, error) {
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

	parseInputFile := func(filePath string) error {
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

			left = append(left, num1)
			right = append(right, num2)
		}

		if len(left) != len(left) {
			return fmt.Errorf("left length (%d) does not match right length (%d)", len(left), len(right))
		}

		return nil
	}

	calcTotalDistance := func(left, right []int) {
		sort.Ints(left)
		sort.Ints(right)

		for i := range len(left) {
			a := left[i]
			b := right[i]

			totalDistance += int(math.Abs(float64(a) - float64(b)))
		}

		fmt.Printf("Total Distance: %d\n", totalDistance)
	}

	calcSimilarityScore := func(left, right []int) {
		for _, num := range right {
			rightOccurences[num]++
		}

		for _, num := range left {
			similarityScore += num * rightOccurences[num]
		}

		fmt.Printf("Similarity Score: %d\n", similarityScore)
	}

	err := parseInputFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	calcTotalDistance(left, right)
	calcSimilarityScore(left, right)
}
