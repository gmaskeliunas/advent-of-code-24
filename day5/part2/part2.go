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
	scanner := bufio.NewScanner(f)
	var rules [][]int
	var updates [][]int
	for scanner.Scan() {

		line := scanner.Text()
		if strings.Contains(line, "|") {
			slice := strings.Split(line, "|")
			var tmp []int
			for _, num := range slice {
				n, _ := strconv.Atoi(num)
				tmp = append(tmp, n)
			}
			rules = append(rules, tmp)
		} else if strings.Contains(line, ",") {
			slice := strings.Split(line, ",")
			var tmp []int
			for _, num := range slice {
				n, _ := strconv.Atoi(num)
				tmp = append(tmp, n)
			}
			updates = append(updates, tmp)
		}

	}

	var correctedUpdates [][]int
	for _, update := range updates {
		var relevantRules [][]int
		for _, rule := range rules {
			if subset(rule, update) {
				relevantRules = append(relevantRules, rule)
			}
		}
		if !correct(update, relevantRules) {
			for !correct(update, relevantRules) {
				fixUpdate(update, relevantRules)
			}
			correctedUpdates = append(correctedUpdates, update)
		}

	}

	sum := 0

	for _, correctedUpdate := range correctedUpdates {
		index := ((len(correctedUpdate) + 1) / 2) - 1
		sum += correctedUpdate[index]
	}

	fmt.Println(sum)
}

func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func correct(update []int, relevantRules [][]int) bool {
	for _, relevantRule := range relevantRules {
		if indexOf(relevantRule[0], update) > indexOf(relevantRule[1], update) {
			return false
		}
	}
	return true
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func fixUpdate(update []int, relevantRules [][]int) []int {
	for _, relevantRule := range relevantRules {
		if indexOf(relevantRule[0], update) > indexOf(relevantRule[1], update) {
			n1 := update[indexOf(relevantRule[0], update)]
			n2 := update[indexOf(relevantRule[1], update)]
			update[indexOf(relevantRule[0], update)] = n2
			update[indexOf(relevantRule[1], update)] = n1
		}
	}
	return update
}
