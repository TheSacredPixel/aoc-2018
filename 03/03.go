package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	length = 1200
)

var (
	claims = 0
	rects  = make([][5]int, 0)
)

func main() {
	reg := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	xMap := [length][length]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //iterate over lines
		//parse rectangle and convert to int
		parsed := reg.FindStringSubmatch(scanner.Text())[1:]
		rect := [5]int{}
		for i, val := range parsed {
			rect[i], _ = strconv.Atoi(val)
		}
		rects = append(rects, rect)
		//calculate area and add to arrays
		for y := rect[1]; y < rect[1]+rect[3]; y++ {
			for x := rect[2]; x < rect[2]+rect[4]; x++ {
				if xMap[x][y]++; xMap[x][y] == 2 {
					claims++
				}
			}
		}

	}
	//print
	fmt.Println(claims)

	//part 2
loop:
	for _, rect := range rects {
		for y := rect[1]; y < rect[1]+rect[3]; y++ {
			for x := rect[2]; x < rect[2]+rect[4]; x++ {
				if xMap[x][y] != 1 {
					continue loop
				}
			}
		}
		fmt.Println(rect[0])
		return
	}
}

//extra func to output the entire area and all claims in a file
func out(x [][]int) {
	out := []byte{}
	for _, line := range x {
		for _, y := range line {
			char := byte(' ')
			switch y {
			case 0:
				char = byte('.')
			case 1:
				char = byte('#')
			default:
				char = []byte(strconv.Itoa(y))[0]
			}
			out = append(out, char)
		}
		out = append(out, '\n')
	}
	file, err := os.Create("out.txt")
	defer file.Close()
	_, err = file.Write(out)
	if err != nil {
		panic(err)
	}
}
