package main

import (
	"fmt"
)

var row int = 0
var column int = 1
var finnish bool = false

var maze = [10][10]string{
	{"*", "S", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", " ", "*", "*", " ", " ", " ", " ", " ", "*"},
	{"*", " ", "*", "*", " ", "*", "*", "*", " ", "*"},
	{"*", " ", "*", "*", "*", "*", "*", "*", " ", "*"},
	{"*", " ", " ", " ", "*", "*", "*", "*", " ", "*"},
	{"*", " ", "*", " ", "*", "*", "*", "*", " ", "*"},
	{"*", " ", "*", " ", "*", " ", " ", " ", " ", "*"},
	{"*", " ", "*", " ", "*", " ", "*", "*", " ", "*"},
	{"*", " ", "*", " ", " ", " ", " ", "*", " ", "F"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
}

func printMaze() {
	for _, i := range maze {
		fmt.Println(i)
	}
	println("---------------------")
}

func move(row, column int) int {
	if finnish == true {
		return 1
	}
	if maze[row+1][column] == " " { // Down
		maze[row+1][column] = "V"
		move(row+1, column)
	} else if maze[row-1][column] == " " { // Up
		maze[row-1][column] = "V"
		move(row-1, column)
	} else if maze[row][column+1] == " " { // Right
		maze[row][column+1] = "V"
		move(row, column+1)
	} else if maze[row][column-1] == " " { // Left
		maze[row][column-1] = "V"
		move(row, column-1)
	} else if maze[row][column+1] == "F" || maze[row][column-1] == "F" || maze[row+1][column] == "F" || maze[row-1][column] == "F" {
		finnish = true
		fmt.Println("Finish")
	} else {
		return 1
	}

	return move(row, column)
}

func main() {
	printMaze()
	move(row, column)
	printMaze()
	return
}
