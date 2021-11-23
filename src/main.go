package main

import (
	"fmt"
	"strconv"
)

const player1 = 0
const player2 = 1
const GRID_SIZE = 9

func main() {
	var grid [GRID_SIZE]int
	var turn bool = false // true = O false = X
	var playing bool = true
	fmt.Println("this is tic tac go!\na simple tic tac toe game writteng in go\nyou can press 'q' to quit!")
	for playing {
		turn = !turn
		for i := 0; i < GRID_SIZE; i++ {
			fmt.Print(" ", grid[i])
			if (i+1)%3 == 0 {
				fmt.Print("\n")
			}
		}
		var input string
		fmt.Print("it is ")
		if turn {
			fmt.Print("O's")
		} else {
			fmt.Print("X's")
		}
		fmt.Print(" turn! ")
		fmt.Scanln(&input)
		if input == "q" {
			playing = false
		} else {
			result, error := strconv.Atoi(input)
			if error == nil {
				fmt.Println("invalid number!")
				continue // return to start of loop
			}

			if result >= 0 && result <= 9 {
				var value int
				if turn {
					value = player1
				} else {
					value = player2
				}
				grid[result-1] = value
			}
		}
	}

}
