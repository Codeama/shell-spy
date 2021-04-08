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

// sorry no American English, not today >.<
type Colour color.Attribute

const (
	Cyan   Colour = 46
	Yellow Colour = 43
	Red    Colour = 101
	Green  Colour = 42
	White  Colour = 107
)

type Option func(*session) error

type session struct {
	TimestampMode bool
	ShellPrompt   string
	Colour        Colour
	Input         io.Reader
	Recorder      io.Writer
	Terminal      io.Writer
}

func WithTimestamps() Option {
	return func(s *session) error {
		s.TimestampMode = true
		return nil
	}
}

func WithUserPrompt(userPrompt string) Option {
	return func(s *session) error {
		s.ShellPrompt = userPrompt
		return nil
	}
}

func WithTerminalColour(userColour Colour) Option {
	return func(s *session) error {
		switch userColour {
		case Yellow:
			s.Colour = Yellow
		case Red:
			s.Colour = Red
		case Green:
			s.Colour = Green
		case White:
			s.Colour = White
		default:
			s.Colour = Cyan
		}
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
func NewSession(filepath string, opts ...Option) (session, error) {
	s := session{}
	f, err := os.Create(filepath)
	if err != nil {
		return session{}, err
	}

	s.Input = os.Stdin
	s.Terminal = os.Stdout
	s.Recorder = f

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

	user, err := user.Current()
	if err != nil {
		return fmt.Errorf("retrieving current user returned err: %v", err)
	}

	host, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("retrieving host name returned err: %v", err)
	}

	prompt := fmt.Sprintf("%v@%v:~$", user.Username, host)

	if s.ShellPrompt != "" {
		prompt = s.ShellPrompt + ":~$"
	}

	if s.TimestampMode {
		s.RecordTime(time.Now())
	}

	if s.Colour == 0 {
		s.Colour = Cyan
	}

	color.New(color.Attribute(s.Colour)).Print(prompt)

	for scanner.Scan() {
		s.Execute(scanner.Text())
		color.New(color.Attribute(s.Colour)).Println(s.Terminal)
	}
	return nil
}

func (s *session) RecordTime(ts time.Time) {
	formattedTime := ts.Format(time.RFC3339)
	fmt.Fprintln(s.Recorder, formattedTime)
}
