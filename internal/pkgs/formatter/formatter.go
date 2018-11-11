package formatter

import (
	"bytes"
	"io"
	"os/exec"
)

// Format returns result of gofmt command.
func Format(src string) (string, error) {
	cmd := exec.Command("gofmt")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	io.WriteString(stdin, src)
	stdin.Close()

	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}
