package lib

import "testing"

func TestRecordUnmarshal(t *testing.T) {
	jsonExamples := []string{
		`{
			"__CURSOR": "s=e9475ee1083b4c42948d18321f8aXXX;i=2736000;b=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXe4d2e375;m=XXXXXX;t=XXXXXX;x=11111111",
			"__REALTIME_TIMESTAMP": "1549905421313842",
			"__MONOTONIC_TIMESTAMP": "1131763146983",
			"_BOOT_ID": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXe4d2e375",
			"_TRANSPORT": "syslog",
			"PRIORITY": "6",
			"SYSLOG_FACILITY": "10",
			"SYSLOG_IDENTIFIER": "CRON",
			"_UID": "0",
			"_GID": "0",
			"_COMM": "cron",
			"_EXE": "/usr/sbin/cron",
			"_CMDLINE": "/usr/sbin/CRON -f",
			"_CAP_EFFECTIVE": "3fffffffff",
			"_SELINUX_CONTEXT": "unconfined\n",
			"_AUDIT_LOGINUID": "0",
			"_SYSTEMD_CGROUP": "/system.slice/cron.service",
			"_SYSTEMD_UNIT": "cron.service",
			"_SYSTEMD_SLICE": "system.slice",
			"_SYSTEMD_INVOCATION_ID": "0a1b200e2d3e455XXXXXXXXXXXXXXXXXXX",
			"_MACHINE_ID": "d560d92f987b4fbfb9CXXXXXXXXXXXXXXXXXXXXXX",
			"_HOSTNAME": "thinkpad",
			"MESSAGE": "pam_unix(cron:session): session closed for user root",
			"SYSLOG_PID": "28301",
			"_PID": "28301",
			"_AUDIT_SESSION": "2891",
			"_SOURCE_REALTIME_TIMESTAMP": "1549905421313798"
		}`,
		`{
			"__CURSOR": "s=d4a113249e9341e29b835XXXXXXXXXX;i=1cef4;b=074c9830ba774b72a3cXXXXXXX;m=XXXXXX;t=XXXXXX;x=11111111",
			"__REALTIME_TIMESTAMP": "1549905531912502",
			"__MONOTONIC_TIMESTAMP": "16004502090",
			"_BOOT_ID": "074c9830ba774b72a3cXXXXXXX",
			"PRIORITY": "6",
			"_UID": "0",
			"_GID": "0",
			"_SYSTEMD_SLICE": "system.slice",
			"_MACHINE_ID": "5083d4a267a54ef59afXXXXXXXXXX",
			"_HOSTNAME": "1.1.1.something.com",
			"_TRANSPORT": "syslog",
			"_CAP_EFFECTIVE": "3fffffffff",
			"SYSLOG_FACILITY": "1",
			"SYSLOG_IDENTIFIER": "example",
			"_COMM": "example",
			"_EXE": "/usr/local/bin/example",
			"_CMDLINE": "/usr/local/bin/example -c /usr/local/etc/example.conf",
			"_SYSTEMD_CGROUP": "/system.slice/example.service",
			"_SYSTEMD_UNIT": "example.service",
			"SYSLOG_PID": "29230",
			"_PID": "29230",
			"MESSAGE": "basket: 1 target: site_sale user: 000000 command: confirmsale status: 200 IP: 1.1.1.1 elapsed: 180.237521ms len: 1792 path: /testing.php",
			"_SOURCE_REALTIME_TIMESTAMP": "1549905531912319"
		}`,
		`{
			"__REALTIME_TIMESTAMP": "1549915794947064",
			"__MONOTONIC_TIMESTAMP": "1142136780204",
			"_BOOT_ID": "XXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"_TRANSPORT": "stdout",
			"PRIORITY": "6",
			"_PID": "2129",
			"_UID": "1000",
			"_GID": "1000",
			"_COMM": "gnome-shell",
			"_EXE": "/usr/bin/gnome-shell",
			"_CMDLINE": "/usr/bin/gnome-shell",
			"_CAP_EFFECTIVE": "0",
			"_SELINUX_CONTEXT": "unconfined\n",
			"_SYSTEMD_CGROUP": "/user.slice/user-1000.slice/session-c2.scope",
			"_SYSTEMD_SESSION": "c2",
			"_SYSTEMD_OWNER_UID": "1000",
			"_SYSTEMD_UNIT": "session-c2.scope",
			"_SYSTEMD_SLICE": "user-1000.slice",
			"_SYSTEMD_USER_SLICE": "-.slice",
			"_SYSTEMD_INVOCATION_ID": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"_MACHINE_ID": "XXXXXXXXXXXXXXXXXXXXXXXXXX",
			"_HOSTNAME": "gabriel-thinkpad",
			"SYSLOG_IDENTIFIER": "firefox.desktop",
			"_STREAM_ID": "7cbe7c3683f146029c42b0124defae12",
			"MESSAGE": "[Parent 1509, Gecko_IOThread] WARNING: pipe error (101): -------------------: file /build/firefox-nSunSF/firefox-99/ipc/chromium/src/chrome/common/ipc_channel_posix.cc, line 122"
		}`,
		`{
			"__REALTIME_TIMESTAMP": "1549918152372677",
			"__MONOTONIC_TIMESTAMP": "1144494205817",
			"_BOOT_ID": "cc3a580059c545f799f9b773e4d2e375",
			"_UID": "1000",
			"_GID": "1000",
			"_CAP_EFFECTIVE": "0",
			"_SELINUX_CONTEXT": "unconfined\n",
			"_SYSTEMD_OWNER_UID": "1000",
			"_SYSTEMD_SLICE": "user-1000.slice",
			"_SYSTEMD_USER_SLICE": "-.slice",
			"_MACHINE_ID": "d560d92f987b4fbfb9a601b73f475d46",
			"_HOSTNAME": "XXXXXX-thinkpad",
			"_AUDIT_SESSION": "2",
			"_AUDIT_LOGINUID": "1000",
			"_SYSTEMD_CGROUP": "/user.slice/user-1000.slice/user@1000.service/dbus.service",
			"_SYSTEMD_UNIT": "user@1000.service",
			"_SYSTEMD_USER_UNIT": "dbus.service",
			"_SYSTEMD_INVOCATION_ID": "XXXXXXXXXXXXXXXXXXXX",
			"PRIORITY": "4",
			"_TRANSPORT": "journal",
			"GLIB_OLD_LOG_API": "1",
			"GLIB_DOMAIN": "GLib",
			"MESSAGE": "g_variant_new_string: assertion 'string != NULL' failed",
			"_COMM": "gnome-clocks",
			"_EXE": "/usr/bin/gnome-clocks",
			"_CMDLINE": "/usr/bin/gnome-clocks --gapplication-service",
			"_PID": "26820",
			"_SOURCE_REALTIME_TIMESTAMP": "1549918152372634"
		}`,
	}

	for _, example := range jsonExamples {
		r := Record{}
		if err := r.UnmarshalJSON([]byte(example)); err != nil {
			t.Errorf("Unmarshal json: %s", err)
			continue
		}

		if r.SOURCE_REALTIME_TIMESTAMP == "" {
			t.Errorf("Invalid SOURCE_REALTIME_TIMESTAMP")
			continue
		}

		if r.MESSAGE == "" {
			t.Errorf("Invalid MESSAGE")
			continue
		}
		t.Logf("OK: %+v", r)
	}
}
