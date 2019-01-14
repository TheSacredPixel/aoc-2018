package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	gridSize = 1000
)

var (
	coords = []map[string]int{}
	xMap   = [gridSize][gridSize]int{}
)

type point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //iterate over lines
		xy := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coords = append(coords, map[string]int{"x": x, "y": y})
	}
	file.Close()

	//if a coordinate is the closest to any point around the edges of our grid,
	//then it must have an infinite area and thus can be excluded,
	//so we hug the wall and scan along the edges
	cleanCoords := make([]map[string]int, len(coords))
	copy(cleanCoords, coords)
	for a, b := range map[string]string{"x": "y", "y": "x"} {
		for _, j := range []int{0, gridSize} {
			for i := 0; i < gridSize; i++ {

				_, found := closestCoord(coords, map[string]int{a: i, b: j})
				if ok, pos := contains(cleanCoords, found); ok {
					slic := cleanCoords[:pos]
					cleanCoords = append(slic, cleanCoords[pos+1:]...)
				}

			}
		}
	}

	//get number of closest points per coordinate
	coordRatings := make(map[int]int)
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			pos, found := closestCoord(coords, map[string]int{"x": x, "y": y})
			if ok, _ := contains(cleanCoords, found); ok && pos != -1 {
				coordRatings[pos]++
			}
			xMap[x][y] = pos
		}
	}

	//get most points
	points, coord := 0, 0
	for c, p := range coordRatings {
		if p > points {
			points = p
			coord = c
		}
	}

	fmt.Println(coordRatings)
	fmt.Printf("%d, at coord ID %d\n", points, coord)
}

func closestCoord(coords []map[string]int, point map[string]int) (int, map[string]int) {
	smallestDist, pos := 999, 0
	closest := map[string]int{}
	for i, coord := range coords {
		dist := map[string]int{"x": 0, "y": 0}
		for _, j := range []string{"x", "y"} { //get dists in x and y
			dist[j] = coord[j] - point[j]
			if dist[j] < 0 {
				dist[j] = -dist[j]
			}
		}
		totalDist := dist["x"] + dist["y"]

		//invalid if 2 dists equal
		if totalDist == smallestDist {
			return -1, nil
		} else if totalDist < smallestDist {
			smallestDist = totalDist
			closest = coord
			pos = i
		}
	}
	//return closest
	return pos, closest
}

func contains(arr []map[string]int, f map[string]int) (bool, int) {
	for i, a := range arr {
		if a["x"] == f["x"] && a["y"] == f["y"] {
			return true, i
		}
	}
	return false, -1
}
