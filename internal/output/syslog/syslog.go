package syslog

import (
	"io"
	"log/syslog"
	"net"
)

type SyslogOutput struct {
	pWC io.WriteCloser
}

func New(c Config) (s *SyslogOutput, err error) {
	priority := getPriority(c.Facility, c.Severity)
	sysLog, err := syslog.Dial(c.Network, net.JoinHostPort(c.Host, c.Port), priority, c.Tag)
	if err != nil {
		return nil, err
	}
	s = &SyslogOutput{
		pWC: sysLog,
	}
	return s, nil
}

func (s *SyslogOutput) Write(b []byte) (n int, err error) {
	return s.pWC.Write(b)
}

func (s *SyslogOutput) Close() error {
	return s.pWC.Close()
}

func getPriority(facility string, severity string) syslog.Priority {
	var f, s syslog.Priority

	switch severity {
	case "LOG_EMERG":
		s = syslog.LOG_EMERG
	case "LOG_ALERT":
		s = syslog.LOG_ALERT
	case "LOG_CRIT":
		s = syslog.LOG_CRIT
	case "LOG_ERR":
		s = syslog.LOG_ERR
	case "LOG_WARNING":
		s = syslog.LOG_WARNING
	case "LOG_NOTICE":
		s = syslog.LOG_NOTICE
	case "LOG_INFO":
		s = syslog.LOG_INFO
	case "LOG_DEBUG":
		s = syslog.LOG_DEBUG
	default:
		s = syslog.LOG_EMERG
	}

	switch facility {
	case "LOG_KERN":
		f = syslog.LOG_KERN
	case "LOG_USER":
		f = syslog.LOG_USER
	case "LOG_MAIL":
		f = syslog.LOG_MAIL
	case "LOG_DAEMON":
		f = syslog.LOG_DAEMON
	case "LOG_AUTH":
		f = syslog.LOG_AUTH
	case "LOG_SYSLOG":
		f = syslog.LOG_SYSLOG
	case "LOG_LPR":
		f = syslog.LOG_LPR
	case "LOG_NEWS":
		f = syslog.LOG_NEWS
	case "LOG_UUCP":
		f = syslog.LOG_UUCP
	case "LOG_CRON":
		f = syslog.LOG_CRON
	case "LOG_AUTHPRIV":
		f = syslog.LOG_AUTHPRIV
	case "LOG_FTP":
		f = syslog.LOG_FTP
	case "LOG_LOCAL0":
		f = syslog.LOG_LOCAL0
	case "LOG_LOCAL1":
		f = syslog.LOG_LOCAL1
	case "LOG_LOCAL2":
		f = syslog.LOG_LOCAL2
	case "LOG_LOCAL3":
		f = syslog.LOG_LOCAL3
	case "LOG_LOCAL4":
		f = syslog.LOG_LOCAL4
	case "LOG_LOCAL5":
		f = syslog.LOG_LOCAL5
	case "LOG_LOCAL6":
		f = syslog.LOG_LOCAL6
	case "LOG_LOCAL7":
		f = syslog.LOG_LOCAL7
	default:
		f = syslog.LOG_KERN
	}

	return s | f
}