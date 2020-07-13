package lib

type Record struct {
	__CURSOR                  string `json:"__CURSOR"`
	REALTIME_TIMESTAMP        string `json:"__REALTIME_TIMESTAMP"`
	__MONOTONIC_TIMESTAMP     string
	_BOOT_ID                  string
	PRIORITY                  string
	_UID                      string `json:"_UID"`
	_GID                      string `json:"_GID"`
	_SYSTEMD_SLICE            string
	_MACHINE_ID               string
	HOSTNAME                  string `json:"_HOSTNAME"`
	_TRANSPORT                string
	_CAP_EFFECTIVE            string
	SYSLOG_FACILITY           string
	SYSLOG_IDENTIFIER         string
	_COMM                     string `json:"_COMM"`
	_EXE                      string `json:"_EXE"`
	_CMDLINE                  string `json:"_CMDLINE"`
	_SYSTEMD_CGROUP           string `json:"_SYSTEMD_CGROUP"`
	_SYSTEMD_UNIT             string `json:"_SYSTEMD_UNIT"`
	SYSLOG_PID                string `json:"SYSLOG_PID"`
	PID                       string `json:"_PID"`
	MESSAGE                   string
	SOURCE_REALTIME_TIMESTAMP string `json:"_SOURCE_REALTIME_TIMESTAMP"`
}
