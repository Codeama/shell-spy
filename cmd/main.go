package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"spy"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	f, err := os.Create("/tmp/log")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var out bytes.Buffer
	for scanner.Scan() {
		spy.Execute(&out, f, scanner.Text())

		fmt.Println(out.String())
		// os.Exit(0)
	}

}
