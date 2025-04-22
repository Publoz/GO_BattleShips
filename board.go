package main

import (
	"fmt"
	"strconv"
)

type Board struct {
	Tiles     [][]*Location
	IsPlayer1 bool
}

// Constructor function for Board
func NewBoard(rows, cols, playerNum int) Board {
	// Initialize a 2D slice of Locations

	if rows <= 3 || cols <= 3 || rows >= 10 || cols >= 10 {
		panic("Invalid board size")
	}

	tiles := make([][]*Location, rows)
	for i := range tiles {
		tiles[i] = make([]*Location, cols)
		for j := range tiles[i] {
			tiles[i][j] = &Location{} // Allocate memory for each Location
		}
	}

	if playerNum == 1 {
		return Board{Tiles: tiles, IsPlayer1: true}
	} else {
		return Board{Tiles: tiles, IsPlayer1: false}
	}

}

func (board Board) AreAllShipsHit() bool {
	for _, row := range board.Tiles {
		for _, tile := range row {
			if tile.HasShip && tile.Status != BattleShipHit {
				return false // A ship exists that hasn't been hit
			}
		}
	}
	return true // All ships have been hit
}

func (board Board) DrawBoard() {
	//rows := len(board.Tiles)
	cols := len(board.Tiles[0])

	if board.IsPlayer1 {
		fmt.Printf(" Player 1's Board\n")
	} else {
		fmt.Println("----------------------")
		fmt.Printf(" Player 2's Board\n")
	}

	// Print column headers
	fmt.Print("  ")
	for col := 0; col < cols; col++ {
		fmt.Print(" " + string('A'+col) + " ")
	}
	fmt.Println()

	// Loop through rows
	for rowNum, row := range board.Tiles {
		// Print row number at the start
		fmt.Print(strconv.Itoa(rowNum+1) + " ")

		// Loop through columns
		for _, location := range row {
			location.DrawLocation()
		}

		// Print row number at the end
		fmt.Println(" " + strconv.Itoa(rowNum+1))
	}

	// Print column headers again at the bottom
	fmt.Print("  ")
	for col := 0; col < cols; col++ {
		fmt.Print(" " + string('A'+col) + " ")
	}
	fmt.Println()
}

// Have to check if tile already has ship
func (b *Board) PlaceBattleship(startRow, startCol int, orientation string) error {
	// Validate starting position
	if startRow < 0 || startCol < 0 || startRow >= len(b.Tiles) || startCol >= len(b.Tiles[0]) {
		return fmt.Errorf("starting position out of bounds")
	}

	// Validate placement based on orientation
	if orientation == "right" {
		if startCol+3 > len(b.Tiles[0]) {
			return fmt.Errorf("battleship does not fit right")
		}
		// Place the battleship horizontally
		for i := 0; i < 3; i++ {
			b.Tiles[startRow][startCol+i].HasShip = true
		}
	} else if orientation == "down" {
		if startRow+3 > len(b.Tiles) {
			return fmt.Errorf("battleship does not fit down")
		}
		// Place the battleship vertically
		for i := 0; i < 3; i++ {
			b.Tiles[startRow+i][startCol].HasShip = true
		}
	} else if orientation == "up" {
		if startRow-2 < 0 {
			return fmt.Errorf("battleship does not fit up")
		}
		// Place the battleship vertically upwards
		for i := 0; i < 3; i++ {
			b.Tiles[startRow-i][startCol].HasShip = true
		}
	} else if orientation == "left" {
		if startCol-2 < 0 {
			return fmt.Errorf("battleship does not fit left")
		}
		// Place the battleship horizontally to the left
		for i := 0; i < 3; i++ {
			b.Tiles[startRow][startCol-i].HasShip = true
		}
	} else {
		return fmt.Errorf("invalid orientation: must be 'right', 'down', 'up', or 'left'")
	}

	return nil
}

func (b Board) IsValidLocation(row, col int) bool {
	return row >= 0 && row < len(b.Tiles) && col >= 0 && col < len(b.Tiles[0])
}
