package spy

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"

	"bitbucket.org/creachadair/shell"
	"github.com/fatih/color"
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
func Execute(terminal io.Writer, recorder io.Writer, commandLine string) error {
	name, args, err := ParseCommand(commandLine)

	w := io.MultiWriter(terminal, recorder)
	fmt.Fprintln(recorder, "************ START ************\nCOMMAND:", commandLine, "\n\nOUTPUT:")

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

func StartSession(filepath string) {
	scanner := bufio.NewScanner(os.Stdin)
	f, err := os.Create(filepath)
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
		Execute(&out, f, scanner.Text())
		color.New(color.FgCyan).Println(out.String())
	}
}
