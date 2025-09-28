package sftp

import (
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	client *sftp.Client
}

func NewSFTPClient(sshClient *ssh.Client) (*Client, error) {
	sc, err := sftp.NewClient(sshClient)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		client: sc,
	}, nil
}

func (s *Client) ReadDir(path string) ([]os.FileInfo, error) {
	return s.client.ReadDir(path)
}

func (s *Client) Open(path string) (*sftp.File, error) {
	return s.client.Open(path)
}

func (s *Client) Create(path string) (*sftp.File, error) {
	return s.client.Create(path)
}

func (s *Client) Remove(path string) error {
	return s.client.Remove(path)
}

func (s *Client) RemoveDirectory(path string) error {
	return s.client.RemoveDirectory(path)
}

func (s *Client) Mkdir(path string) error {
	return s.client.Mkdir(path)
}

func (s *Client) GetWd() (string, error) {
	return s.client.Getwd()
}

func (s *Client) Stat(path string) (os.FileInfo, error) {
	return s.client.Stat(path)
}

func (s *Client) Close() error {
	return s.client.Close()
}
