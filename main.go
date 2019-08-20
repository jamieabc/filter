package main

import (
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
	info, err := os.Stdin.Stat()
	if nil != err {
		panic(err)
	}

	if !isCharDevice(info.Mode()) {
		fmt.Println("The command is intended to work with pipes.")
		return
	}

	timer := time.After(initialCheckTimeSecond)

	for {
		select {
		case <-timer:
			_, err := io.Copy(os.Stdout, os.Stdin)
			if nil != err {
				panic(err)
			}
			timer = time.After(checkTimeSecond)
		}
	}
}

func isCharDevice(mode os.FileMode) bool {
	return 0 == mode&os.ModeCharDevice
}
