package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	player1Board, player2Board := handleSetup()

	playerTurn := 1

	// Main game loop
	for !player1Board.AreAllShipsHit() && !player2Board.AreAllShipsHit() {
		player1Board.DrawBoard()
		player2Board.DrawBoard()

		fmt.Println("###############")

		var boardAttacking *Board
		if playerTurn == 1 {
			boardAttacking = player2Board
		} else {
			boardAttacking = player1Board
		}

		fmt.Printf("Player %d, Please fire! (row,col) e.g 2,3\n", playerTurn)
		row, col := parseLocation()
		if row == -1 || col == -1 || !boardAttacking.IsValidLocation(row, col) {
			fmt.Println("Invalid numbers. Please enter valid input.")
			continue
		}

		var locationFired *Location
		locationFired = boardAttacking.Tiles[row][col]

		locationFired.Shoot()

		if playerTurn == 1 {
			playerTurn = 2
		} else {
			playerTurn = 1
		}
	}

	fmt.Println("Game Over!")

}

func handleSetup() (*Board, *Board) {
	fmt.Print("Select battleships to play with: ")
	var ships int
	fmt.Scan(&ships)

	fmt.Print("Select Grid size to play with: ")
	var gridSize int
	fmt.Scan(&gridSize)

	player1Board := makeBoard(1, gridSize, ships)
	player2Board := makeBoard(2, gridSize, ships)

	return player1Board, player2Board

}

func parseLocation() (int, int) {
	var input string
	fmt.Scan(&input)

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		fmt.Println("Invalid input. Please enter in the format row,col")
		return -1, -1
	}

	row, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	col, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers. Please enter valid integers.")
		return -1, -1
	}

	return row, col
}

func makeBoard(playerNum int, gridSize int, ships int) *Board {
	fmt.Printf("Player %d, please place ships. Comma separate (row, col) e.g 3,2", playerNum)
	playerBoard := NewBoard(gridSize, gridSize, playerNum)

	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			playerBoard.Tiles[row][col] = &Location{Row: row, Col: col, Status: Unknown, HasShip: false}
		}
	}

	for i := 0; i < ships; i++ {
		fmt.Printf("\nEnter position for ship %d: ", i+1)

		row, col := parseLocation()

		if row == -1 || col == -1 {
			fmt.Println("Invalid numbers. Please enter valid input.")
			i-- // Retry this iteration
			continue
		}

		location := Location{Row: row, Col: col, Status: Unknown, HasShip: true}
		validPlaces := location.ValidPlacementLocations(gridSize)
		if validPlaces == nil || !playerBoard.IsValidLocation(row, col) {
			fmt.Println("Invalid placement. Please try again.")
			i-- // Retry this iteration
			continue
		}

		fmt.Printf("Valid placements are: %v\n", validPlaces)
		var input string
		fmt.Scan(&input)

		// Check if the input direction is in validPlaces
		isValid := false
		for _, direction := range validPlaces {
			if direction == input {
				isValid = true
				break
			}
		}

		if !isValid {
			fmt.Println("Invalid direction. Please try again.")
			i-- // Retry this iteration
			continue
		}

		playerBoard.PlaceBattleship(row, col, input)

	}

	return &playerBoard

}
