package main

import (
	"math/rand"
)

const (
	boardWidth  = 128
	boardHeight = 48
)

const (
	alive = "#"
	dead  = " "
)

var cellTypeArray = [...]string{alive, dead}

type boardType [boardHeight][boardWidth]string

func newBoard() *boardType {
	b := &boardType{}

	for rowI := 0; rowI < boardHeight; rowI++ {
		for colI := 0; colI < boardWidth; colI++ {
			cell := cellTypeArray[rand.Intn(len(cellTypeArray))]

			b[rowI][colI] = cell
		}
	}

	return b
}

func (b *boardType) String() string { // big S so it will consider in the fmt.Stringer interface
	var s string

	for _, row := range b {
		for _, cell := range row {
			s += cell
		}
		s += "\n"
	}
	return s
}

func cellAlive(cell string) bool {
	return cell == alive
}

func (b *boardType) update() {
	oldB := *b

	for rowI := 0; rowI < boardHeight; rowI++ {
		for colI := 0; colI < boardWidth; colI++ {
			aliveNeighbors := 0

			for rowOffset := -1; rowOffset <= 1; rowOffset++ {
				for colOffset := -1; colOffset <= 1; colOffset++ {
					if rowOffset == 0 && colOffset == 0 {
						continue // its the original spot should continue
					}

					newRowI := rowI + rowOffset
					newColI := colI + colOffset

					if newRowI < 0 || newColI < 0 || newRowI >= boardHeight || newColI >= boardWidth {
						continue // its off bounds no such cell
					}

					checkingCell := oldB[newRowI][newColI]
					if cellAlive(checkingCell) {
						aliveNeighbors++
					}
				}
			}

			cell := oldB[rowI][colI]

			if cellAlive(cell) {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					// should die
					b[rowI][colI] = dead
				}
			} else {
				if aliveNeighbors == 3 {
					// should become alive
					b[rowI][colI] = alive
				}
			}
		}
	}
}

// github copilots implementation
/*
func (b *boardType) update() {
	oldB := *b

	for rowI := 0; rowI < boardHeight; rowI++ {
		for colI := 0; colI < boardWidth; colI++ {
			neighbors := 0

			for rowOffset := -1; rowOffset <= 1; rowOffset++ {
				for colOffset := -1; colOffset <= 1; colOffset++ {
					if rowOffset == 0 && colOffset == 0 {
						continue
					}

					neighborRowI := rowI + rowOffset
					neighborColI := colI + colOffset

					if neighborRowI < 0 || neighborRowI >= boardHeight {
						continue
					}

					if neighborColI < 0 || neighborColI >= boardWidth {
						continue
					}

					neighbor := oldB[neighborRowI][neighborColI]

					if neighbor == alive {
						neighbors++
					}
				}
			}

			if oldB[rowI][colI] == alive {
				if neighbors < 2 || neighbors > 3 {
					b[rowI][colI] = dead
				}
			} else {
				if neighbors == 3 {
					b[rowI][colI] = alive
				}
			}
		}
	}
}
*/
