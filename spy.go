package spy

import (
	"bufio"
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
func (s *Session) Execute(commandLine string) error {
	name, args, err := ParseCommand(commandLine)

	w := io.MultiWriter(s.Terminal, s.Recorder)
	fmt.Fprintln(s.Recorder, "COMMAND:", commandLine, "\n\nOUTPUT:")

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

type Session struct {
	timestampMode bool
	prompt        bool
	Recorder      io.Writer
	Terminal      io.Writer
}

func WithTimestamps() Option {
	return func(s *Session) error {
		s.timestampMode = true
		return nil
	}
}

func WithCustomPrompt() Option {
	return func(s *Session) error {
		s.prompt = true
		return nil
	}
}

// NewSession creates and returns a new shell session with a file
// and specified options to customise the terminal
func NewSession(filepath string, opts ...Option) (Session, error) {
	s := Session{}
	f, err := os.Create(filepath)
	if err != nil {
		return Session{}, err
	}
	s.Terminal = os.Stdout
	s.Recorder = f

	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return Session{}, err
		}
	}
	return s, nil
}

func (s *Session) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	if s.timestampMode {
		s.RecordTime(time.Now())
	}
	color.New(color.BgCyan).Printf("%v@%v:", user.Username, host)
	for scanner.Scan() {
		s.Execute(scanner.Text())
		color.New(color.FgCyan).Println(s.Terminal)
	}
}
func (s *Session) RecordTime(ts time.Time) {
	formattedTime := ts.Format(time.RFC3339)
	fmt.Fprintf(s.Recorder, "Log time: %s\n\n", formattedTime)

}
