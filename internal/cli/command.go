package cli

import (
	"context"
)

// ICommand represents a CLI command that can be executed
type ICommand interface {
	Execute(ctx context.Context, session *SessionContext, args []string) error
	Name() string
	Help() string
	Usage() string
}

// ICommandRegistry manages available command
type ICommandRegistry interface {
	Register(cmd ICommand)
	Get(name string) (ICommand, bool)
	List() []ICommand
	Execute(ctx context.Context, name string, session *SessionContext, args []string) error
}
