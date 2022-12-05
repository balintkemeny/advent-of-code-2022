package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ErrInvalidItemRune error = errors.New("invalid item rune")
	ErrNoDuplicateItem error = errors.New("no duplicate item found")
	ErrNoCommonItem    error = errors.New("no common item found in group")
	ErrGrouping        error = errors.New("number of total rucksacks is not divisible by 3")
)

func main() {
	inputFilePath := "input.txt"
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", inputFilePath)
	}

	fmt.Println(ex01(input))
	fmt.Println(ex02(input))
}

func ex01(input []byte) int {
	var sumOfPriorities int

	rucksacks := getRucksacksCompartmentalised(input)
	for _, rucksack := range rucksacks {
		duplicateItem, err := getDuplicateItemInRucksack(rucksack)
		if err != nil {
			log.Fatalf("Error: %s for rucksack: %v", err.Error(), rucksack)
		}

		duplicateItemPriority, err := determineItemPriority(duplicateItem)
		if err != nil {
			log.Fatalf("Error: %s for item: %c", err.Error(), duplicateItem)
		}

		sumOfPriorities += duplicateItemPriority
	}

	return sumOfPriorities
}

func ex02(input []byte) int {
	var sumOfCommonItemPriorities int

	rucksacks := getRucksacksTotal(input)
	groups, err := formGroups(rucksacks)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	for _, group := range groups {
		commonItem, err := getCommonItemInGroup(group)
		if err != nil {
			log.Fatalf("Error: %s for group: %v", err.Error(), group)
		}

		commonItemPriority, err := determineItemPriority(commonItem)
		if err != nil {
			log.Fatalf("Error: %s for item: %c", err.Error(), commonItem)
		}

		sumOfCommonItemPriorities += commonItemPriority
	}

	return sumOfCommonItemPriorities
}

func getRucksacksCompartmentalised(input []byte) [][2]string {
	var rucksacks [][2]string
	for _, row := range bytes.Split(input, []byte{'\n'}) {
		rowStr := string(row)
		splitPoint := len(row) / 2

		rucksack := [2]string{rowStr[:splitPoint], rowStr[splitPoint:]}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func getRucksacksTotal(input []byte) []string {
	var rucksacks []string
	for _, row := range bytes.Split(input, []byte{'\n'}) {
		rucksacks = append(rucksacks, string(row))
	}

	return rucksacks
}

func getDuplicateItemInRucksack(rucksack [2]string) (rune, error) {
	compartmentA := rucksack[0]
	compartmentB := rucksack[1]

	for _, item := range compartmentA {
		itemFound := strings.IndexRune(compartmentB, item) != -1
		if itemFound {
			return item, nil
		}
	}

	return '0', ErrNoDuplicateItem
}

func determineItemPriority(item rune) (int, error) {
	lowercaseOffset := -96
	uppercaseOffset := -38

	itemAsciiValue := int(item)
	if itemAsciiValue >= 65 && itemAsciiValue <= 90 {
		return itemAsciiValue + uppercaseOffset, nil
	}

	if itemAsciiValue >= 97 && itemAsciiValue <= 122 {
		return itemAsciiValue + lowercaseOffset, nil
	}

	return -1, ErrInvalidItemRune
}

func formGroups(rucksacks []string) ([][3]string, error) {
	rucksackCount := len(rucksacks)
	if rucksackCount%3 != 0 {
		return nil, ErrGrouping
	}

	groupCount := rucksackCount / 3
	groups := make([][3]string, groupCount)

	var i int
	for _, rucksack := range rucksacks {
		elf := i % 3
		group := i / 3
		groups[group][elf] = rucksack
		i++
	}

	return groups, nil
}

func getCommonItemInGroup(group [3]string) (rune, error) {
	for _, item := range group[0] {
		foundInSecond := strings.IndexRune(group[1], item) != -1
		foundInThird := strings.IndexRune(group[2], item) != -1

		if foundInSecond && foundInThird {
			return item, nil
		}
	}

	return '0', ErrNoCommonItem
}
