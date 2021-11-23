package main

import (
	"fmt"
	"strconv"
)

const GRID_SIDE = 3
const GRID_SIZE = 9

func clear_console() {
	fmt.Print("\033[H\033[2J")
}

func did_win(grid [GRID_SIZE]string) bool {

	for i := 0; i < GRID_SIDE; i++ {
		/*
			this looks through all the vertical win possibilities
		*/
		if grid[i] == grid[i+GRID_SIDE] && grid[i+GRID_SIDE] == grid[i+GRID_SIDE*2] && grid[i] != " " {
			return true
		}
		/*
			this looks through all the horizontal win possibilities
		*/
		if grid[i*GRID_SIDE] == grid[i*GRID_SIDE+1] && grid[i*GRID_SIDE+1] == grid[i*GRID_SIDE+2] && grid[i*GRID_SIDE] != " " {
			return true
		}
	}

	// the diagonals
	if grid[0] == grid[4] && grid[4] == grid[8] && grid[0] != " " {
		return true
	}
	if grid[2] == grid[4] && grid[4] == grid[6] && grid[2] != " " {
		return true
	}

	return false
}

func get_player_turn(turn bool) string {
	if turn {
		return "O"
	}
	return "X"
}

func is_draw(grid [GRID_SIZE]string) bool {
	for i := 0; i < GRID_SIZE; i++ {
		if grid[i] == " " {
			return false
		}
	}
	return true
}

func main() {
	var grid [GRID_SIZE]string
	for i := 0; i < GRID_SIZE; i++ {
		grid[i] = " "
	}
	turn := false // true = O false = X
	playing := true
	fmt.Print("this is tic tac go!\na simple tic tac toe game writteng in go\nyou can press 'q' to quit!\n\n")
	for playing {
		turn = !turn // invert turn
		// print grid
		for j := 0; j < 7; j++ {
			fmt.Print("─")
		}
		fmt.Print("\n")
		for i := 0; i < GRID_SIZE; i++ {
			fmt.Print("|")
			fmt.Print(grid[i])
			if (i+1)%3 == 0 {
				fmt.Print("|")
				fmt.Print("\n")
				for j := 0; j < 7; j++ {
					fmt.Print("─")
				}
				fmt.Print("\n")
			}
		}
		var input string
		fmt.Print("\nit is ", get_player_turn(turn), "'s turn!", " type a number between 1-9: ")
		fmt.Scanln(&input)
		if input == "q" {
			playing = false
		} else {
			result, error := strconv.Atoi(input)
			if error != nil {
				fmt.Println("invalid number!")
				continue // return to start of loop
			}

			if result >= 0 && result <= 9 {
				value := "O"
				if !turn {
					value = "X"
				}
				result--
				if grid[result] == " " {
					grid[result] = value
				}
			}
		}

		clear_console()
		if is_draw(grid) {
			playing = false
			fmt.Println("Draw!")
		} else if did_win(grid) {
			fmt.Println(get_player_turn(turn), "Won the game!")
			playing = false
		}
	}

}
