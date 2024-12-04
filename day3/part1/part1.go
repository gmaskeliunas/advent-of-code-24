package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(f)

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	allMuls := r.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, mul := range allMuls {
		d1, _ := strconv.Atoi(mul[1])
		d2, _ := strconv.Atoi(mul[2])
		sum += d1 * d2
	}

	fmt.Println(sum)
}
