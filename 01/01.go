package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sum  = 0
	freq = make([]int, 0)
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		sum += i
		freq = append(freq, i)
	}
	fmt.Println(sum)

	//part 2
	sum2 := 0
	seen := make(map[int]bool)
	for {
		for _, i := range freq {
			if seen[sum2] {
				fmt.Println(sum2)
				return
			}
			seen[sum2] = true
			sum2 += i
		}
	}
}
