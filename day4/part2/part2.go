package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	var xmasLines [][]string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()
		slice := strings.Split(line, "")
		for i := 0; i < 3; i++ {
			slice = append(slice, "0")
		}
		xmasLines = append(xmasLines, slice)
	}
	for i := 0; i < 3; i++ {
		xmasLines = append(xmasLines, make([]string, len(xmasLines)+3))
	}
	count := 0

	for i := 0; i < len(xmasLines)-3; i++ {
		for j := 0; j < len(xmasLines[i])-3; j++ {
			if xmasLines[i][j] == "M" && xmasLines[i+1][j+1] == "A" && xmasLines[i+2][j+2] == "S" && xmasLines[i][j+2] == "M" && xmasLines[i+2][j] == "S" {
				count++
			}
			if xmasLines[i][j] == "M" && xmasLines[i+1][j+1] == "A" && xmasLines[i+2][j+2] == "S" && xmasLines[i][j+2] == "S" && xmasLines[i+2][j] == "M" {
				count++
			}
			if xmasLines[i][j] == "S" && xmasLines[i+1][j+1] == "A" && xmasLines[i+2][j+2] == "M" && xmasLines[i][j+2] == "M" && xmasLines[i+2][j] == "S" {
				count++
			}
			if xmasLines[i][j] == "S" && xmasLines[i+1][j+1] == "A" && xmasLines[i+2][j+2] == "M" && xmasLines[i][j+2] == "S" && xmasLines[i+2][j] == "M" {
				count++
			} else {
				continue
			}
		}
	}
	fmt.Println(count)
}
