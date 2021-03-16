package spy

import (
	"fmt"
	"io"
	"os/exec"

	"bitbucket.org/creachadair/shell"
)

// ParseCommand accepts a string and returns a command name and a slice of args
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

// Execute accepts two writers and a shell command
// It writes to both writers; whilst the first writer logs only
// the executed command result on the user's terminal
// the second logs to a specified writer with additional messages
func Execute(w io.Writer, w2 io.Writer, commandLine string) error {
	name, args, err := ParseCommand(commandLine)
	m := io.MultiWriter(w, w2)

	w2.Write([]byte("************ START ************\n"))
	w2.Write([]byte(fmt.Sprintf("COMMAND: %s\n\nOUTPUT:\n", commandLine)))

	if err != nil {
		fmt.Fprintln(m, "Error parsing command")
		return nil
	}
	cmd := exec.Command(name, args...)
	cmd.Stdout = m

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
