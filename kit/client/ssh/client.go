package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	opts Options
}

func New(opts Options) *Client {
	return &Client{opts: opts}
}

func (c *Client) Dial() (*ssh.Client, error) {
	if c.opts.Host == "" {
		return nil, ErrHostEmpty
	}

	if c.opts.Username == "" {
		return nil, ErrUserEmpty
	}

	if c.opts.Port == 0 {
		c.opts.Port = 22
	}

	address := fmt.Sprintf("%v:%d", c.opts.Host, c.opts.Port)
	cfg := c.toClientConfig()
	return ssh.Dial("tcp", address, cfg)
}

func (c *Client) toClientConfig() *ssh.ClientConfig {
	var auth []ssh.AuthMethod
	auth = append(auth, ssh.Password(c.opts.Password))

	var config ssh.Config
	extrasCiphers := c.opts.CiphersExtras

	if len(extrasCiphers) > 0 {
		supportedCiphers := ssh.SupportedAlgorithms().Ciphers
		config.Ciphers = rmDuplicates(supportedCiphers, extrasCiphers)
	}

	extrasKEX := c.opts.KeyExchangesExtras
	if len(extrasKEX) > 0 {
		supportedKEX := ssh.SupportedAlgorithms().KeyExchanges
		config.KeyExchanges = rmDuplicates(supportedKEX, extrasKEX)
	}

	extrasMACs := c.opts.MACsExtras
	if len(extrasMACs) > 0 {
		supportedMACs := ssh.SupportedAlgorithms().MACs
		config.MACs = rmDuplicates(supportedMACs, extrasMACs)
	}

	return &ssh.ClientConfig{
		User:            c.opts.Username,
		Auth:            auth,
		HostKeyCallback: c.resolveHostKeyCallback(),
		Timeout:         c.opts.Timeout,
		Config:          config,
	}
}

func (c *Client) resolveHostKeyCallback() ssh.HostKeyCallback {
	if c.opts.IgnoreHostKeyCheck {
		return ssh.InsecureIgnoreHostKey()
	}
	if c.opts.KnownHostsCallback != nil {
		return c.opts.KnownHostsCallback
	}
	return ssh.InsecureIgnoreHostKey()
}

func rmDuplicates[T comparable](slices ...[]T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, slice := range slices {
		for _, item := range slice {
			if !seen[item] {
				result = append(result, item)
				seen[item] = true
			}
		}
	}

	return result
}
