package systemdJournald

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gabrielperezs/sytoco/lib"
)

const (
	defaultCmd = "journalctl --utc -l -f -n 5 -o json"
)

var (
	newLine           = []byte("\n")
	newLineR          = []byte("\r")
	removeWhiteSpaces = regexp.MustCompile(`\s+`)
)

type Config struct {
	Cmd string
}

type SystemdJournald struct {
	running *exec.Cmd
	cmd     string
	ctx     context.Context
	cancel  context.CancelFunc
	outputs []lib.Output
	buf     *bytes.Buffer
}

func New(cfg *Config, outputs ...lib.Output) *SystemdJournald {
	sj := &SystemdJournald{
		outputs: outputs,
		buf:     bytes.NewBuffer(make([]byte, 0, 1024)),
	}
	sj.config(cfg)
	go sj.listener()
	return sj
}

func (sj *SystemdJournald) config(cfg *Config) {
	sj.cmd = defaultCmd

	if cfg == nil {
		return
	}

	if cfg.Cmd != "" {
		sj.cmd = string(removeWhiteSpaces.ReplaceAll([]byte(cfg.Cmd), []byte(" ")))
	}
}

func (sj *SystemdJournald) listener() error {
	sj.ctx, sj.cancel = context.WithCancel(context.Background())
	args := strings.Split(sj.cmd, " ")
	sj.running = exec.CommandContext(sj.ctx, args[0], args[1:]...)
	sj.running.Stdout = sj
	if err := sj.running.Run(); err != nil {
		log.Panicf("sytoco/SystemdJournald ERROR cmd: %s", err)
		return err
	}
	return nil
}

func (sj *SystemdJournald) Write(p []byte) (n int, err error) {
	// If is not a full json message we just write in the
	// loca buffer and continue
	if !bytes.HasSuffix(p, newLine) {
		return sj.buf.Write(p)
	}

	// Write in the buffer
	sj.buf.Write(p)

	scanner := bufio.NewScanner(sj.buf)
	for scanner.Scan() {
		r := lib.Record{}
		if err := r.UnmarshalJSON(scanner.Bytes()); err != nil {
			//log.Printf("sytoco/SystemdJournald Unmarshal error: %s <| %s |>", err, scanner.Text())
			continue
		}
		for _, o := range sj.outputs {
			o.Write(r)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("sytoco/SystemdJournald ERROR scanner: %s", err)
	}
	return len(p), nil
}

func (sj *SystemdJournald) Exit() {
	go sj.cancel()
	<-sj.ctx.Done()
}
