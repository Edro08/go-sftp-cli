package command

import (
	"context"
	"go-sftp-cli/internal/cli"
)

// ClearCommand implements the clear command
type ClearCommand struct{}

func NewClearCommand() cli.ICommand {
	return &ClearCommand{}
}

func (c *ClearCommand) Name() string {
	return "clear"
}

func (c *ClearCommand) Help() string {
	return "Limpia la pantalla"
}

func (c *ClearCommand) Usage() string {
	return "clear"
}

func (c *ClearCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	// Clear screen using ANSI escape codes
	session.GetUI().Print("\033[2J\033[H")
	return nil
}
