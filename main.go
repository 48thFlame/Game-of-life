package main

import (
	"fmt"
	"math/rand"
	"time"
)

const timeDelay = time.Millisecond * 100

func init() {
	rand.Seed(time.Now().UnixNano())

	initClearMap()
}

func main() {
	b := newBoard()
	for {
		b.update()
		fmt.Println(b)
		time.Sleep(timeDelay)
		callClear()
	}
}
