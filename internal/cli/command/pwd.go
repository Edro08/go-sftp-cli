package command

import (
	"context"
	"go-sftp-cli/internal/cli"
)

// PwdCommand implements the pwd command
type PwdCommand struct{}

func NewPwdCommand() cli.ICommand {
	return &PwdCommand{}
}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) Help() string {
	return "Muestra el directorio actual"
}

func (c *PwdCommand) Usage() string {
	return "pwd"
}

func (c *PwdCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	session.GetUI().Println(session.GetCurrentDir())
	return nil
}
