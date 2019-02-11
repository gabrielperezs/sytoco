package systemdJournald

import (
	"testing"
)

func TestConfigDefault(t *testing.T) {
	sj := &SystemdJournald{}
	sj.config(nil)

	if sj.cmd != defaultCmd {
		t.Errorf("It's not reading the default configuration")
	}
}

func TestConfigCmdCustom(t *testing.T) {
	sj := &SystemdJournald{}
	sj.config(&Config{
		Cmd: "/usr/sbin/journalctl -f -n 0 -o json",
	})

	if sj.cmd != "/usr/sbin/journalctl -f -n 0 -o json" {
		t.Errorf("Can't read the custom command")
	}
}

func TestConfigCmdCustomWithSpaces(t *testing.T) {
	sj := &SystemdJournald{}
	sj.config(&Config{
		Cmd: "/usr/sbin/journalctl -f -n 0     -o json",
	})

	if sj.cmd != "/usr/sbin/journalctl -f -n 0 -o json" {
		t.Errorf("It's not removing the white spaces in the custom command")
	}
}
