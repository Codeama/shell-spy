package shell

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

// Parser represents a function that parses string text
// and returns a tring and slice of strings
type Parser func(string) (string, []string)

// ParseCommand accepts a string and returns the command name and args
func ParseCommand(cmd string) (string, []string) {
	stringResult := strings.Split(cmd, " ")
	return stringResult[0], stringResult[1:]
}

// Execute accepts a command argument and writes the result to a writer
func Execute(writer io.Writer, parserFunc Parser, commandLine string) {
	name, args := parserFunc(commandLine)
	cmd := exec.Command(name, args...)
	cmd.Stdout = writer
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(writer)
}
