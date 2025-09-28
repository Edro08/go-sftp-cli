package client

import "golang.org/x/crypto/ssh"

type SSHClient interface {
	Dial() (*ssh.Client, error)
}
