package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	var reports [][]int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		intTemp := make([]int, len(temp))
		for i, s := range temp {
			intTemp[i], _ = strconv.Atoi(s)
		}

		reports = append(reports, intTemp)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, report := range reports {
		if isDecreasing(report) || isIncreasing(report) {
			count++
		}
	}
	fmt.Println(count)
}

func isIncreasing(report []int) bool {

	for i := 0; i < len(report)-1; i++ {
		if !((report[i+1] > report[i]) && (report[i+1]-report[i] >= 1) && (report[i+1]-report[i] <= 3)) {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {

	for i := 0; i < len(report)-1; i++ {
		if !((report[i] > report[i+1]) && (report[i]-report[i+1] >= 1) && (report[i]-report[i+1] <= 3)) {
			return false
		}
	}
	return true
}
