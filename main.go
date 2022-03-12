package main

import (
	"time"
)

func main() {
	for {
		BuildTable()
		time.Sleep(time.Duration(5) * time.Second)
	}
}
