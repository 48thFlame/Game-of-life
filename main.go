package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	boardWidth  = 48
	boardHeight = 12
)

const timeDelay = time.Millisecond * 100

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

var clearScreenMapFunc map[string]func() //create a map for storing clear funcs

func init() {
	rand.Seed(time.Now().UnixNano())

	clearScreenMapFunc = make(map[string]func()) //Initialize it
	clearScreenMapFunc["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearScreenMapFunc["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clearScreenMapFunc[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	for {
		b := newBoard()
		fmt.Println(b)
		time.Sleep(timeDelay)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		// .Run()
		// time.Sleep(timeDelay)
	}
}
