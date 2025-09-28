package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
)

// HelpCommand implements the help command
type HelpCommand struct {
	registry cli.ICommandRegistry
}

func NewHelpCommand(registry cli.ICommandRegistry) cli.ICommand {
	return &HelpCommand{registry: registry}
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Help() string {
	return "Muestra la ayuda de comandos disponibles"
}

func (c *HelpCommand) Usage() string {
	return "help [comando]"
}

func (c *HelpCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	ui := session.GetUI()

	if len(args) > 0 {
		// Show help for specific command
		cmd, exists := c.registry.Get(args[0])
		if !exists {
			ui.Printf("❌ Comando desconocido: %s\n", args[0])
			return fmt.Errorf("comando desconocido: %s", args[0])
		}

		ui.Printf("Comando: %s\n", cmd.Name())
		ui.Printf("Descripción: %s\n", cmd.Help())
		ui.Printf("Uso: %s\n", cmd.Usage())
		return nil
	}

	// Show all command
	ui.Println("\n📋 Comandos disponibles:")
	ui.Println("======================")

	commands := c.registry.List()
	for _, cmd := range commands {
		ui.Printf("%-10s - %s\n", cmd.Name(), cmd.Help())
	}

	ui.Println("\nComandos especiales:")
	ui.Println("exit, quit - Salir del programa")
	ui.Println("\nUsa 'help <comando>' para obtener ayuda específica de un comando.")
	ui.Println("")

	return nil
}
