package main

import (
	"fmt"
	"slices"
)

type position struct {
	x int8
	y int8
}

func VisitHouses(directions string) int8 {
	pos := position{x: 0, y: 0}
	visitedPos := []position{}
    totalHouses := 0;

    for i := 0; i <= len(directions); i++ {
        curPos := position{x: pos.x, y: pos.y}
        if !slices.Contains(visitedPos, curPos) {
            totalHouses += 1;
            visitedPos = append(visitedPos, curPos);
        }

        // skip last iteration
        if i == len(directions) {
            continue;
        }

        char := directions[i];
		switch char {
		case '^':
			pos.y += 1
		case '>':
			pos.x += 1
		case 'v':
			pos.y -= 1
		case '<':
			pos.x -= 1
		}
    }

    return int8(totalHouses);
}

func main() {
	directions := ">"
    result := VisitHouses(directions);

    fmt.Printf("%d\n", result);
}
