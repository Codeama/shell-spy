package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
		w := io.MultiWriter(f, &out)
		spy.Execute(w, scanner.Text())
		if len(scanner.Text()) == 0 {
			out.Reset()
		}
		fmt.Println(out.String())
	}

}
