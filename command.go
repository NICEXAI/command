package command

import (
	"io"
	"log"
	"os/exec"
	"strings"
)

// Cmd represents an external command being prepared or run
type Cmd struct {
	cmd *exec.Cmd
}

func (c *Cmd) Wait() error {
	return c.cmd.Wait()
}

func (c *Cmd) Stop() error {
	return c.cmd.Process.Kill()
}

// Run that can output logs
func Run(name string, arg ...string) (*Cmd, error) {
	var (
		stdout io.ReadCloser
		stderr io.ReadCloser
		err    error
	)

	command := exec.Command(name, arg...)

	stdout, err = command.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err = command.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err = command.Start(); err != nil {
		return nil, err
	}

	go outputLog(stdout)
	go outputLog(stderr)

	return &Cmd{cmd: command}, nil
}

func outputLog(reader io.ReadCloser) {
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		logByte := buf[:num]

		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "closed") {
				err = nil
			}
			break
		}
		if num > 0 {
			log.Println(string(logByte))
		}
	}
}
