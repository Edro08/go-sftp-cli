package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
)

// commandRegistry implements ICommandRegistry interface
type commandRegistry struct {
	commands map[string]cli.ICommand
}

func NewCommandRegistry() cli.ICommandRegistry {
	return &commandRegistry{
		commands: make(map[string]cli.ICommand),
	}
}

func (r *commandRegistry) Register(cmd cli.ICommand) {
	r.commands[cmd.Name()] = cmd
}

func (r *commandRegistry) Get(name string) (cli.ICommand, bool) {
	cmd, exists := r.commands[name]
	return cmd, exists
}

func (r *commandRegistry) List() []cli.ICommand {
	commands := make([]cli.ICommand, 0, len(r.commands))
	for _, cmd := range r.commands {
		commands = append(commands, cmd)
	}
	return commands
}

func (r *commandRegistry) Execute(ctx context.Context, name string, session *cli.SessionContext, args []string) error {
	cmd, exists := r.commands[name]
	if !exists {
		return fmt.Errorf("comando desconocido: %s", name)
	}
	return cmd.Execute(ctx, session, args)
}
