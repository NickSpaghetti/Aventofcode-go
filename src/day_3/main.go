package main

import (
	"fmt"
	"os"
	"strings"
)

const input = "input.txt"
const sample = "sample.txt"

var (
	letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func main() {
	//question1(input)
	question2(input)
}

func question1(f string) {
	rawRuckSack := readRuckSack(f)
	splitRuckSack := strings.Split(rawRuckSack, "\n")
	sum := 0
	for _, s := range splitRuckSack {
		sum += calcSum(s)
		fmt.Println(sum)
	}
}

func question2(f string) {
	rawRuckSack := readRuckSack(f)
	splitRuckSack := strings.Split(rawRuckSack, "\n")
	sum := 0
	for i := 0; i < len(splitRuckSack); i += 3 {
		elfGroups := take(splitRuckSack, i, i+3)
		if elfGroups == nil {
			fmt.Printf("failed take any values for start %d end %d, len(%d) \n", i, i+3, len(splitRuckSack))
			os.Exit(-1)
		}
		result := findDistinctLetter([]rune(elfGroups[0]), []rune(elfGroups[1]), []rune(elfGroups[2]))
		if result == -1 {
			fmt.Printf("failed to find any common letters in \n %s \n %s \n %s\n", elfGroups[0], elfGroups[1], elfGroups[2])
			os.Exit(-1)
		}
		sum += result
	}
	fmt.Println(sum)
}

func readRuckSack(fileName string) string {
	byteSlice, error := os.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	sack := string(byteSlice)
	return sack
}

func calcSum(ruckSackLine string) int {
	first := ruckSackLine[:len(ruckSackLine)/2]
	second := ruckSackLine[len(ruckSackLine)/2:]
	firstMap := map[string]int{}
	secondMap := map[string]int{}
	cntMap := map[string]int{}

	for _, f := range first {
		updateMapCount(firstMap, string(f))
	}
	for _, s := range second {
		updateMapCount(secondMap, string(s))
	}

	for fk, _ := range firstMap {
		_, found := secondMap[fk]
		if found {
			index := indexOfLetter(fk, 1)
			if index == -1 {
				os.Exit(-1)
			}
			cntMap[fk] = index
		}
	}
	maxValue := 0
	for _, v := range cntMap {
		if v >= maxValue {
			maxValue = v
		}
	}

	if maxValue == 0 {
		os.Exit(2)
	}

	return maxValue

}

func updateMapCount(myMap map[string]int, key string) {
	if _, found := myMap[key]; found {
		myMap[key] += 1
	} else {
		myMap[key] = 1
	}
}

func indexOfLetter(letter string, offSet int) int {
	for i, v := range letters {
		if letter == v {
			return i + offSet
		}
	}
	return -1
}

func findDistinctLetter(e []rune, el []rune, elf []rune) int {

	for _, e1 := range e {
		if isContainedIn(el, e1) && isContainedIn(elf, e1) {
			return indexOfLetter(string(e1), 1)
		}
	}
	return -1
}

func isContainedIn(s []rune, l rune) bool {
	for v, _ := range s {
		if l == s[v] {
			return true
		}
	}
	return false
}

func take(lst []string, start int, end int) []string {
	taken := []string{}
	for i := start; i < end; i++ {
		if i >= len(lst) {
			return nil
		} else {
			taken = append(taken, lst[i])
		}

	}
	return taken
}
