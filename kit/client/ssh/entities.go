package ssh

import (
	"errors"
	"time"

	"golang.org/x/crypto/ssh"
)

type Options struct {
	Host               string
	Port               int
	Username           string
	Password           string
	Timeout            time.Duration
	CiphersExtras      []string
	KeyExchangesExtras []string
	MACsExtras         []string
	IgnoreHostKeyCheck bool
	KnownHostsCallback ssh.HostKeyCallback
}

var (
	ErrHostEmpty = errors.New("host cannot be empty")
	ErrUserEmpty = errors.New("user cannot be empty")
)
