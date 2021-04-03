package main

import (
	"log"
	"spy"
)

func main() {
	s, err := spy.NewSession("/tmp/log",
		spy.WithTimestamps(),
	)
	if err != nil {
		log.Fatal(err)
	}
	s.Run()
}
