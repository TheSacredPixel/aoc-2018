package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)

	result := react(input)

	fmt.Println(len(result) - 1)

	//part 2
	length := 0
	for i := 65; i < 91; i++ {
		test := strings.Replace(input, string(i), "", -1)
		test = strings.Replace(test, string(i+32), "", -1)

		result = react(test)
		if length == 0 || len(result) < length {
			length = len(result)
		}
	}

	fmt.Println(length - 1)
}

func react(input string) string {
	for {
		found := false
		for i := 0; i < len(input)-1; i++ {
			//find pair to trim
			if input[i]+32 == input[i+1] || input[i]-32 == input[i+1] {
				//fmt.Printf("trimming %c and %c\n", input[i], input[i+1])
				//trim
				input = input[:i] + input[i+2:]
				found = true
			}
		}
		if !found {
			break
		}
	}
	return input
}
