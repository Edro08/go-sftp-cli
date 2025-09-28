package main

import (
	"go-sftp-cli/internal/cli/command"
	"go-sftp-cli/internal/cli/shell"
)

func Run() {
	registry := command.NewCommandRegistry()
	registry.Register(command.NewHelpCommand(registry))
	registry.Register(command.NewLsCommand())
	registry.Register(command.NewCdCommand())
	registry.Register(command.NewPwdCommand())
	registry.Register(command.NewMkdirCommand())
	registry.Register(command.NewRmdirCommand())
	registry.Register(command.NewRmCommand())
	registry.Register(command.NewGetCommand())
	registry.Register(command.NewPutCommand())
	registry.Register(command.NewStatCommand())
	registry.Register(command.NewClearCommand())
	registry.Register(command.NewCommandAlias("dir", "ls", registry))

	newUI := shell.NewConsoleUI()
	shell.NewInteractive(registry, newUI).Run()
}
