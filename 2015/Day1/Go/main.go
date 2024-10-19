// 2015
// --- Day 1: Not Quite Lisp --- 
package main

import "fmt"

func GetNewFloorFromDirections(floor *int, directions string) {
    for _, char := range directions {
        if char == '(' {
            *floor += 1;
        }

        if char == ')' {
            *floor -= 1;
        }
    }
}

func main()  {
    floor := 0;
    directions := "()()";

    GetNewFloorFromDirections(&floor, directions);

    fmt.Printf("%d\n", floor);
}
