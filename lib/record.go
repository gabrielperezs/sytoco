package lib

type Record struct {
	__CURSOR                  string
	REALTIME_TIMESTAMP        string `json:"__REALTIME_TIMESTAMP"`
	__MONOTONIC_TIMESTAMP     string
	_BOOT_ID                  string
	PRIORITY                  string
	_UID                      string
	_GID                      string
	_SYSTEMD_SLICE            string
	_MACHINE_ID               string
	HOSTNAME                  string `json:"_HOSTNAME"`
	_TRANSPORT                string
	_CAP_EFFECTIVE            string
	SYSLOG_FACILITY           string
	SYSLOG_IDENTIFIER         string
	_COMM                     string
	_EXE                      string
	_CMDLINE                  string
	_SYSTEMD_CGROUP           string
	_SYSTEMD_UNIT             string
	SYSLOG_PID                string
	PID                       string `json:"_PID"`
	MESSAGE                   string
	SOURCE_REALTIME_TIMESTAMP string `json:"_SOURCE_REALTIME_TIMESTAMP"`
}
