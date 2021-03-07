package shell

import (
	"fmt"
	"io"
	"os/exec"

	"bitbucket.org/creachadair/shell"
)

// ParseCommand accepts a string and returns the command name and args
func ParseCommand(cmd string) (string, []string) {
	stringResult, _ := shell.Split(cmd)

	if len(stringResult) <= 1 {
		return cmd, []string{}
	}
	return stringResult[0], stringResult[1:]
}

// Execute accepts a command argument and writes the result to a writer
func Execute(w io.Writer, commandLine string) error {
	name, args := ParseCommand(commandLine)
	cmd := exec.Command(name, args...)
	cmd.Stdout = w
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Print(w)
	return nil
}
