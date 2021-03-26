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
	"time"

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

type Option func(*Session) error

type Session struct{
	timestampMode bool
	Recorder io.Writer
}

func WithTimestamps() Option {
	return func(s *Session) error {
		s.timestampMode = true
		return nil
	}
}

func NewSession(filepath string, opts ...Option) (Session, error) {
	session := Session{}
	f, err := os.Create(s.filepath)
	if err != nil {
		return err
	}
	session.Recorder = f
	for _, opt := range opts {
		err := opt(&session)
		if err != nil {
			return err
		}
	}
	return session
}

func (s *Session) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	defer f.Close()
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	if s.timestampMode {
		// do the timestamps
	}
	var out bytes.Buffer
	color.New(color.BgCyan).Printf("%v@%v:", user.Username, host)
	for scanner.Scan() {
		s.Execute(scanner.Text())
		color.New(color.FgCyan).Println(out.String())
	}
}

func (s *Session) Record(ts time.Time, text string) {
	formattedTime := ts.Format(time.RFC3339)
	fmt.Fprintf(s.Recorder, "%s %s", formattedTime, s.Recorder)
}