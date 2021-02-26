package shell

import (
	"bytes"
	"log"
	"os/exec"
)

// Execute accepts a command argument and returns the result of the executed command
func Execute(script string) string {
	cmd := exec.Command(script)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
