package main

import (
	"fmt"
	"math/rand"
	"time"
)

const timeDelay = time.Millisecond * 25

func init() {
	rand.Seed(time.Now().UnixNano())

	initClearMap()
}

func main() {
	b := newBoard()
	for {
		callClear()
		fmt.Println(b)
		time.Sleep(timeDelay)
		b.update()
	}
}
