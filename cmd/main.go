package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/user"
	"spy"

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	f, err := os.Create("/tmp/log")
	if err != nil {
		log.Fatal(err)
	}

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var out bytes.Buffer
	color.New(color.BgCyan).Printf("%v@%v:", user.Username, host)
	for scanner.Scan() {
		spy.Execute(&out, f, scanner.Text())
		color.New(color.FgCyan).Println(out.String())
	}

}
