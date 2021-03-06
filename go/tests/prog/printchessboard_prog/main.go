package main

import (
	"fmt"
	"os"
	"strconv"
)

func solve(x, y int) {
	for i := 0; i < x; i++ {
		line := ""
		for j := 0; j < y; j++ {
			if i%2 == 0 && j%2 == 0 {
				line += "#"
			} else if i%2 == 0 && j%2 == 1 {
				line += " "
			} else if i%2 == 1 && j%2 == 1 {
				line += "#"
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error")
		return
	}
	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[0])
	if x <= 0 || y <= 0 {
		fmt.Println("Error")
	} else {
		solve(x, y)
	}
}
