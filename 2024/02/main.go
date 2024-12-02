package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var safeReports int

	parseLine := func(line string) ([]int, error) {
		numbers := strings.Fields(line)

		var res []int
		for _, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}

			res = append(res, n)
		}

		return res, nil
	}

	isReportSafe := func(report []int) error {
		type check struct {
			Ascending  bool
			Descending bool
			Initialise bool
		}
		orderCheck := check{Ascending: false, Descending: false, Initialise: false}

		if len(report) <= 1 {
			return fmt.Errorf("report contains no levels")
		}
		for i, num := range report {
			if i == 0 {
				continue
			}

			diff := int(math.Abs(float64(num) - float64(report[i-1])))
			if diff <= 0 || diff > 3 {
				return fmt.Errorf("difference out of range")
			}

			if i == 1 && num > report[i-1] {
				orderCheck.Ascending = true
				orderCheck.Initialise = true
				continue
			}

			if i == 1 && num < report[i-1] {
				orderCheck.Descending = true
				orderCheck.Initialise = true
				continue
			}

			if i == 1 && num == report[i-1] {
				return fmt.Errorf("neither an increase or decrease")
			}

			if orderCheck.Descending && orderCheck.Initialise {
				if num >= report[i-1] {
					return fmt.Errorf("order not correct")
				}
			}

			if orderCheck.Ascending && orderCheck.Initialise {
				if num <= report[i-1] {
					return fmt.Errorf("order not correct")
				}
			}

		}
		return nil
	}

	problemDampener := func(report []int) bool {
		for i := range len(report) {
			newSlice := make([]int, len(report))
			copy(newSlice, report)
			newSlice = append(newSlice[:i], newSlice[i+1:]...)
			err := isReportSafe(newSlice)
			if err != nil {
				continue
			} else {
				return true
			}
		}

		return false
	}

	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report, err := parseLine(line)
		if err != nil {
			log.Panic(err)
		}

		err = isReportSafe(report)
		if err == nil {
			safeReports++
		} else {
			if problemDampener(report) {
				safeReports++
			} else {
				continue
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	fmt.Printf("%d reports are safe\n", safeReports)
}
