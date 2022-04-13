package main

import "math/rand"

const (
	boardWidth  = 48
	boardHeight = 12
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

func (b *boardType) String() string {
	var s string

	for _, row := range b {
		for _, cell := range row {
			s += cell
		}
		s += "\n"
	}
	return s
}
