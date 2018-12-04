package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	tsFormat = "2006-01-02 15:04"
)

type eventType struct {
	timestamp time.Time
	str       string
}

func main() {
	//regTime := regexp.MustCompile(`\[1518-\d+-\d+ \d+:\d+\]`)
	regEvent := regexp.MustCompile(`\d+|asleep|wakes`)

	days := make(map[string][]eventType)
	guardSleepLog := make(map[int]map[int]int) //[guard][min]times
	guardMinsSleeping := make(map[int]int)     //[guard]mins

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //iterate over lines
		//parse timestamp and event
		str := scanner.Text()[1:]
		ind := strings.Index(str, "]")
		ts, event := str[:ind], regEvent.FindStringSubmatch(str[ind+2:])[0]
		timestamp, err := time.Parse(tsFormat, ts)
		if err != nil {
			panic(err)
		}

		//assign times to relevant day
		day := ""
		//catch if shift starts before midnight
		if timestamp.Hour() == 23 {
			day = timestamp.Add(time.Hour * 24).Format("0102")
		} else {
			day = timestamp.Format("0102")
		}
		days[day] = append(days[day], eventType{timestamp, event})
	}

	//determine guard ID and time asleep per day
	for _, events := range days {
		//sort chronologically
		sort.SliceStable(events, func(i, j int) bool {
			return events[i].timestamp.Before(events[j].timestamp)
		})

		guardID, slept := 0, 0
		for _, event := range events {
			switch event.str {
			case "asleep": //sleep event
				slept = event.timestamp.Minute()
			case "wakes": //awake event
				for i := slept; i < event.timestamp.Minute(); i++ {
					if guardSleepLog[guardID] == nil {
						guardSleepLog[guardID] = make(map[int]int)
					}
					guardSleepLog[guardID][i]++
					guardMinsSleeping[guardID]++
				}
			default: //guard shift start
				guardID, err = strconv.Atoi(event.str)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	//find least active guard...
	mostMins, guardID := 0, 0
	for ID, mins := range guardMinsSleeping {
		if mins > mostMins {
			mostMins = mins
			guardID = ID
		}
	}

	//...and guard's minute with most sleeping occurences
	topMin, _ := findTopMin(guardSleepLog[guardID])

	fmt.Println(guardID * topMin)

	//part 2
	//find most slept minute per guard, and lowest from all guards
	targetGuard, targetMin, targetOcc := 0, 0, 0
	for guardID, guard := range guardSleepLog {
		topMin, topOcc := findTopMin(guard)
		if topOcc > targetOcc {
			targetMin = topMin
			targetGuard = guardID
			targetOcc = topOcc
		}
	}

	fmt.Println(targetGuard * targetMin)
}

func findTopMin(guard map[int]int) (int, int) {
	topMin, topOcc := 0, 0
	for min, occ := range guard {
		if occ > topOcc {
			topOcc = occ
			topMin = min
		}
	}
	return topMin, topOcc
}
