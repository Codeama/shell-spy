package main

import "log"

func main() {
	err := spy.RunSession("/tmp/log",
		spy.WithTimestamps(),
	)
	if err != nil {
		log.Fatal(err)
	}
}
