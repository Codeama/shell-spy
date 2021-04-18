package spy

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"time"

	"bitbucket.org/creachadair/shell"
	"github.com/fatih/color"
)

type option func(*session) error

type session struct {
	TimestampMode bool
	ShellPrompt   string
	Colour        color.Attribute
	Input         io.Reader
	Recorder      io.Writer
	Terminal      io.Writer
}

func WithTimestamps() option {
	return func(s *session) error {
		s.TimestampMode = true
		return nil
	}
}

func WithUserPrompt(userPrompt string) option {
	return func(s *session) error {
		s.ShellPrompt = userPrompt + ":~$"
		return nil
	}
}

func WithTerminalColour(userColour color.Attribute) option {
	return func(s *session) error {
		s.Colour = userColour
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
func (s *session) Execute(commandLine string) error {
	name, args, err := ParseCommand(commandLine)

	w := io.MultiWriter(s.Terminal, s.Recorder)
	fmt.Fprintln(s.Recorder, commandLine)

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
func NewSession(filepath string, opts ...option) (session, error) {
	f, err := os.Create(filepath)
	if err != nil {
		return session{}, err
	}

	prompt, err := defaultPrompt()
	if err != nil {
		return session{}, err
	}

	s := session{
		Input:       os.Stdin,
		Terminal:    os.Stdout,
		Recorder:    f,
		ShellPrompt: prompt,
		Colour:      color.BgCyan,
	}

	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return session{}, err
		}
	}

	return s, nil
}

func (s *session) Run() error {
	scanner := bufio.NewScanner(s.Input)
	colour := color.New(s.Colour)

	s.RecordTime(time.Now())

	colour.Print(s.ShellPrompt)

	for scanner.Scan() {
		s.Execute(scanner.Text())
		colour.Println(s.Terminal)
	}
	return nil
}

func (s *session) RecordTime(ts time.Time) {
	if s.TimestampMode {
		formattedTime := ts.Format(time.RFC3339)
		fmt.Fprintln(s.Recorder, formattedTime)
	}
}

func defaultPrompt() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("retrieving current user returned err: %v", err)
	}

	host, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("retrieving host name returned err: %v", err)
	}

	prompt := fmt.Sprintf("%v@%v:~$", user.Username, host)

	return prompt, nil
}
