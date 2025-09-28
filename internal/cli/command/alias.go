package command

import (
	"context"
	"go-sftp-cli/internal/cli"
	"strings"
)

// AliasCommand implements ICommand interface for command aliases
type AliasCommand struct {
	name     string
	target   string
	registry cli.ICommandRegistry
}

func NewCommandAlias(name string, target string, registry cli.ICommandRegistry) *AliasCommand {
	return &AliasCommand{
		name:     name,
		target:   target,
		registry: registry,
	}
}

func (a *AliasCommand) Name() string {
	return a.name
}

func (a *AliasCommand) Help() string {
	if cmd, exists := a.registry.Get(a.target); exists {
		return cmd.Help()
	}
	return "Alias para " + a.target
}

func (a *AliasCommand) Usage() string {
	if cmd, exists := a.registry.Get(a.target); exists {
		return strings.Replace(cmd.Usage(), a.target, a.name, 1)
	}
	return a.name
}

func (a *AliasCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	return a.registry.Execute(ctx, a.target, session, args)
}
