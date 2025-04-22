package main

import "fmt"

const battleShipSize int = 3

type Location struct {
	Row     int
	Col     int
	Status  SquareStatus
	HasShip bool
}

type SquareStatus int

const (
	Unknown       SquareStatus = iota // 0
	BattleShipHit                     // 1
	Miss                              // 2
)

func (status SquareStatus) locationString() string {
	switch status {
	case Unknown:
		return "|_|"
	case BattleShipHit:
		return "|X|"
	case Miss:
		return "|O|"
	default:
		panic("Invalid SquareStatus")
	}
}

func (location *Location) Shoot() {
	if location.HasShip {
		location.Status = BattleShipHit
	} else {
		location.Status = Miss
	}
}

func (loc Location) ValidPlacementLocations(gridSize int) []string {
	validDirections := []string{}

	// Check if placing downwards is valid (Row increases)
	if loc.Row+(battleShipSize-1) < gridSize {
		validDirections = append(validDirections, "down")
	}

	// Check if placing upwards is valid (Row decreases)
	if loc.Row-(battleShipSize-1) >= 0 {
		validDirections = append(validDirections, "up")
	}

	// Check if placing rightwards is valid (Col increases)
	if loc.Col+(battleShipSize-1) < gridSize {
		validDirections = append(validDirections, "right")
	}

	// Check if placing leftwards is valid (Col decreases)
	if loc.Col-(battleShipSize-1) >= 0 {
		validDirections = append(validDirections, "left")
	}

	return validDirections
}

func (location Location) DrawLocation() {
	fmt.Print(location.Status.locationString())
}
