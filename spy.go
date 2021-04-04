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

type Option func(*Session) error

type Session struct {
	TimestampMode bool
	ShellPrompt   string
	Recorder      io.Writer
	Terminal      io.Writer
}

func WithTimestamps() Option {
	return func(s *Session) error {
		s.TimestampMode = true
		return nil
	}
}

func WithUserPrompt(userPrompt string) Option {
	return func(s *Session) error {
		s.ShellPrompt = userPrompt
		return nil
	}
}

// ParseCommand accepts a string and returns a command name and a slice of args
func ParseCommand(cmd string) (string, []string, error) {
	stringResult, ok := shell.Split(cmd)

	if !ok {
		return "", []string{}, fmt.Errorf("cannot parse command: %q", cmd)
	}

	if len(stringResult) <= 1 {
		return cmd, []string{}, nil
	}
	return stringResult[0], stringResult[1:], nil
}

// Execute accepts a shell command and writes the result to
// the Session struct writers (Recorder and Terminal)
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

// NewSession creates and returns a new shell session with a file
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

	prompt := fmt.Sprintf("%v@%v:~$", user.Username, host)

	if s.ShellPrompt != "" {
		prompt = s.ShellPrompt + ":~$"
	}

	if s.TimestampMode {
		s.RecordTime(time.Now())
	}

	color.New(color.BgCyan).Printf(prompt)
	for scanner.Scan() {
		s.Execute(scanner.Text())
		color.New(color.FgCyan).Println(s.Terminal)
	}
}

func (s *Session) RecordTime(ts time.Time) {
	formattedTime := ts.Format(time.RFC3339)
	fmt.Fprintf(s.Recorder, "Log time: %s\n\n", formattedTime)
}
