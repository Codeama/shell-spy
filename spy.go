package spy

import (
	"fmt"
	"io"
	"os/exec"

	"bitbucket.org/creachadair/shell"
)

// ParseCommand accepts a string and returns the command name and args
func ParseCommand(cmd string) (string, []string, error) {
	stringResult, ok := shell.Split(cmd)

	if !ok {
		return "", []string{}, fmt.Errorf("Cannot parse command: %q", cmd)
	}

	if len(stringResult) <= 1 {
		return cmd, []string{}, nil
	}
	return stringResult[0], stringResult[1:], nil
}

// Execute accepts a command argument and writes the result to a writer
func Execute(w io.Writer, commandLine string) error {
	name, args, err := ParseCommand(commandLine)
	fmt.Fprintln(w, commandLine)

	if err != nil {
		fmt.Fprintln(w, "Error parsing command")
		return nil
	}
	cmd := exec.Command(name, args...)

	cmd.Stdout = w

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
