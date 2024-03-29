package main

import (
	"filter/arguments"
	"fmt"
	"io"
	"os"

	"time"
)

const (
	initialCheckTimeSecond = 200 * time.Millisecond
	checkTimeSecond        = 5 * time.Second
)

func main() {
	keyword := arguments.Parse()
	if "" == keyword {
		printHelp()
		return
	}

	info, err := os.Stdin.Stat()
	if nil != err {
		panic(err)
	}

	if !isCharDevice(info.Mode()) {
		fmt.Println("The command is intended to work with pipes.")
		return
	}

	timer := time.NewTimer(initialCheckTimeSecond)

	for {
		select {
		case <-timer.C:
			_, err := io.Copy(os.Stdout, os.Stdin)
			if nil != err {
				panic(err)
			}
			timer.Reset(checkTimeSecond)
		}
	}
}

func printHelp() {
	fmt.Printf("Usage: filter -keyword=[search word]\n")
	fmt.Println("Please input search keyword")
}

func isCharDevice(mode os.FileMode) bool {
	return 0 == mode&os.ModeCharDevice
}
