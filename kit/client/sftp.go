package client

import (
	"github.com/pkg/sftp"
	"os"
)

type SFTPClient interface {
	ReadDir(path string) ([]os.FileInfo, error)
	Open(path string) (*sftp.File, error)
	Create(path string) (*sftp.File, error)
	Remove(path string) error
	RemoveDirectory(path string) error
	Mkdir(path string) error
	GetWd() (string, error)
	Stat(path string) (os.FileInfo, error)
	Close() error
}
