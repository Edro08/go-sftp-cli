package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
)

// RmCommand implements the rm command
type RmCommand struct{}

func NewRmCommand() cli.ICommand {
	return &RmCommand{}
}

func (c *RmCommand) Name() string {
	return "rm"
}

func (c *RmCommand) Help() string {
	return "Elimina un archivo"
}

func (c *RmCommand) Usage() string {
	return "rm <archivo>"
}

func (c *RmCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: rm <archivo>")
		return fmt.Errorf("archivo requerido")
	}

	client := session.GetSFTPClient()
	ui := session.GetUI()

	for _, file := range args {
		err := client.Remove(file)
		if err != nil {
			ui.Printf("❌ Error al eliminar archivo %s: %v\n", file, err)
			return err
		}
		ui.Printf("✅ Archivo eliminado: %s\n", file)
	}
	return nil
}
