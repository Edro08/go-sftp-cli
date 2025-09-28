package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
)

// MkdirCommand implements the mkdir command
type MkdirCommand struct{}

func NewMkdirCommand() cli.ICommand {
	return &MkdirCommand{}
}

func (c *MkdirCommand) Name() string {
	return "mkdir"
}

func (c *MkdirCommand) Help() string {
	return "Crea un directorio"
}

func (c *MkdirCommand) Usage() string {
	return "mkdir <directorio>"
}

func (c *MkdirCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: mkdir <directorio>")
		return fmt.Errorf("directorio requerido")
	}

	client := session.GetSFTPClient()
	ui := session.GetUI()

	for _, dir := range args {
		err := client.Mkdir(dir)
		if err != nil {
			ui.Printf("❌ Error al crear directorio %s: %v\n", dir, err)
			return err
		}
		ui.Printf("✅ Directorio creado: %s\n", dir)
	}
	return nil
}
