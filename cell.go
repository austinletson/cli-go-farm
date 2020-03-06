package main

import (
	"strings"
)

type CellType int

const (
	PLAYER CellType = iota
	POINTER
	GRASS
	WATER
)

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

const worldY = 4
const worldX = 5

var world = [4][5]CellType{
	{GRASS, GRASS, WATER, GRASS, GRASS},
	{GRASS, GRASS, WATER, WATER, GRASS},
	{GRASS, GRASS, GRASS, GRASS, GRASS},
	{WATER, GRASS, GRASS, GRASS, GRASS},
}

var cellPointer = [2]int{0, 0}

func getWorldMatrix() [worldY][worldX]CellType {
	return world
}

func getWorldMatrixWithPointer() [worldY][worldX]CellType {
	worldCopy := world
	worldCopy[cellPointer[0]][cellPointer[1]] = POINTER
	return worldCopy
}

func getWorldString() string {
	var strs []string
	for _, row := range getWorldMatrixWithPointer() {
		rowString := []string{}

		for i := range row {
			number := row[i]
			text := getCellGraphic(number)
			rowString = append(rowString, text)
		}
		s := strings.Join(rowString, "")
		strs = append(strs, s)

	}
	s := strings.Join(strs, "\n")
	return s
}

func drawCellPointerOnWorld() {
}

func drawCellsOnWorld(cells [][]CellType, locX int, locY int) {
	for y, row := range cells {
		for x, cell := range row {
			world[y+locY][x+locX] = cell
		}
	}
}

func getCellGraphic(cellType CellType) string {

	cellGraphics := make(map[CellType]string)

	cellGraphics[POINTER] = "+"
	cellGraphics[GRASS] = "G"
	cellGraphics[WATER] = "W"

	cellTypeString := cellGraphics[cellType]
	return cellTypeString
}

type Cell struct {
	graphic  string
	cellType CellType
}

func moveCellPointer(d Direction) {
	if d == NORTH && pointInBounds(cellPointer[0]-1, cellPointer[1]) {
		cellPointer[0]--
	} else if d == SOUTH && pointInBounds(cellPointer[0]+1, cellPointer[1]) {
		cellPointer[0]++
	} else if d == EAST && pointInBounds(cellPointer[0], cellPointer[1]+1) {
		cellPointer[1]++
	} else if d == WEST && pointInBounds(cellPointer[0], cellPointer[1]-1) {
		cellPointer[1]--
	}
}

func pointInBounds(y int, x int) bool {
	return x < worldX && x >= 0 && y < worldY && y >= 0
}
