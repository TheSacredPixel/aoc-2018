package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	twos   = 0
	threes = 0
	inputs = make([]string, 0)
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //iterate over lines
		chars := make(map[rune]int)
		seenTwo, seenThree := false, false
		inputs = append(inputs, scanner.Text())

		//get occurences per char
		for _, char := range scanner.Text() {
			chars[char]++
		}

		//check for 2- or 3-pairs, adding only if seen for the first time
		for _, times := range chars {
			if times == 2 && !seenTwo {
				twos++
				seenTwo = true
			}
			if times == 3 && !seenThree {
				threes++
				seenThree = true
			}
		}
	}

	fmt.Println(twos * threes)

	//part 2
	for pos, str1 := range inputs {//first str to compare
		for i := pos + 1; i < len(inputs); i++ {//second str to compare
			diffs := make([]int, 0)

			for charPos, char := range str1 {//go over 1st str's chars
				if char != []rune(inputs[i])[charPos] {//check for differences with 2nd str
					diffs = append(diffs, charPos)
				}
			}

			if len(diffs) == 1 {
				fmt.Println(str1[:diffs[0]] + str1[diffs[0]+1:])//remove different char and print
				return
			}
		}
	}
}
