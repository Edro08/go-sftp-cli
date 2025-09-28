package cli

import "go-sftp-cli/kit/client/sftp"

type SessionContext struct {
	currentDir string
	sftpClient *sftp.Client
	ui         IUserInterface
}

// NewSessionContext creates a new session context
func NewSessionContext(sftpClient *sftp.Client, ui IUserInterface) *SessionContext {
	return &SessionContext{
		currentDir: "/",
		sftpClient: sftpClient,
		ui:         ui,
	}
}

// GetCurrentDir returns the current directory of the session
func (s *SessionContext) GetCurrentDir() string {
	return s.currentDir
}

// SetCurrentDir sets the current directory of the session
func (s *SessionContext) SetCurrentDir(dir string) {
	if dir == "" {
		dir = "/"
	}
	s.currentDir = dir
}

// GetSFTPClient returns the SFTP client of the session
func (s *SessionContext) GetSFTPClient() *sftp.Client {
	return s.sftpClient
}

// GetUI returns the user interface of the session
func (s *SessionContext) GetUI() IUserInterface {
	return s.ui
}
