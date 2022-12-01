package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("cannot open input file")
	}

	var elves []int
	var maxElf int

	for _, elf := range bytes.Split(input, []byte("\n\n")) {
		var elfTotal int

		for _, foodItemRaw := range bytes.Split(elf, []byte("\n")) {
			foodItem, err := strconv.Atoi(string(foodItemRaw))
			if err != nil {
				log.Fatalf("Cannot convert to integer: %s", foodItemRaw)
			}

			elfTotal += foodItem
		}

		elves = append(elves, elfTotal)
		if elfTotal > maxElf {
			maxElf = elfTotal
		}
	}

	fmt.Printf("The elf carrying the most energy carries: %d calories.\n", maxElf)

	sort.Ints(elves)
	var topThreeTotal int
	for i := 1; i <= 3; i++ {
		topThreeTotal += elves[len(elves)-i]
	}

	fmt.Printf("The total calories carried by the top 3 elves is: %d \n", topThreeTotal)
}
