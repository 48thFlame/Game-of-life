package main

import (
	"os"
	"os/exec"
	"runtime"
)

var clearScreenMapFunc map[string]func() //create a map for storing clear funcs

func initClearMap() {
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

func callClear() {
	value, ok := clearScreenMapFunc[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                                       //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
