package shell

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

// ParseCommand accepts a string and returns the command name and args
func ParseCommand(cmd string) (string, []string) {
	stringResult := strings.Split(cmd, " ")
	return stringResult[0], stringResult[1:]
}

// Execute accepts a command argument and writes the result to a writer
func Execute(writer io.Writer, parser func(string) (string, []string), commandLine string) {
	name, args := parser(commandLine)
	cmd := exec.Command(name, args...)
	cmd.Stdout = writer
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(writer)
}
