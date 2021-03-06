package main

import (
	"log"
	"spy"
)

func main() {
	s, err := spy.NewSession("/tmp/log",
		spy.WithTimestamps(),
		spy.WithUserPrompt("Type here"),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
