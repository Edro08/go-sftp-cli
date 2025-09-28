package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
)

// RmdirCommand implements the rmdir command
type RmdirCommand struct{}

func NewRmdirCommand() cli.ICommand {
	return &RmdirCommand{}
}

func (c *RmdirCommand) Name() string {
	return "rmdir"
}

func (c *RmdirCommand) Help() string {
	return "Elimina un directorio vacío"
}

func (c *RmdirCommand) Usage() string {
	return "rmdir <directorio>"
}

func (c *RmdirCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: rmdir <directorio>")
		return fmt.Errorf("directorio requerido")
	}

	client := session.GetSFTPClient()
	ui := session.GetUI()

	for _, dir := range args {
		err := client.RemoveDirectory(dir)
		if err != nil {
			ui.Printf("❌ Error al eliminar directorio %s: %v\n", dir, err)
			return err
		}
		ui.Printf("✅ Directorio eliminado: %s\n", dir)
	}
	return nil
}
