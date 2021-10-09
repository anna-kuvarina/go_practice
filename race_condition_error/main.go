package main

import "time"

var (
	sharedInt = 0
	unusedInt = 0
)

func startSimpleReader() {
	for {
		value := sharedInt
		if value%10 == 0 {
			unusedInt++
		}
	}
}

func startSimpleWriter() {
	for {
		sharedInt++
	}
}

func main() {
	go startSimpleReader()
	go startSimpleWriter()
	time.Sleep(10 * time.Second)
}
