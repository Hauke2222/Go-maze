package main

import "fmt"

var row int = 0
var column int = 1

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

func move(row, column *int) int {

	if maze[*row+1][*column] != "*" { // Down
		maze[*row][*column] = "V"
		maze[*row+1][*column] = "S"
		*row += 1
	} else if maze[*row-1][*column] != "*" { //	Up
		maze[*row][*column] = "V"
		maze[*row-1][*column] = "S"
		*row -= 1
	} else if maze[*row][*column+1] != "*" { // Right
		maze[*row][*column] = "V"
		maze[*row][*column+1] = "S"
		*column += 1
	} else if maze[*row][*column-1] != "*" { // Left
		maze[*row][*column] = "V"
		maze[*row][*column-1] = "S"
		*column -= 1
	} else {
		return 1
	}

	printMaze()
	// return 1
	return move(row, column)
}

func main() {
	printMaze()
	move(&row, &column)
	return
}
