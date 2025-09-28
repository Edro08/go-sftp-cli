package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
)

// StatCommand implements the stat command
type StatCommand struct{}

func NewStatCommand() cli.ICommand {
	return &StatCommand{}
}

func (c *StatCommand) Name() string {
	return "stat"
}

func (c *StatCommand) Help() string {
	return "Muestra información detallada de un archivo o directorio"
}

func (c *StatCommand) Usage() string {
	return "stat <archivo>"
}

func (c *StatCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: stat <archivo>")
		return fmt.Errorf("archivo requerido")
	}

	client := session.GetSFTPClient()
	ui := session.GetUI()

	for _, path := range args {
		stat, err := client.Stat(path)
		if err != nil {
			ui.Printf("❌ Error al obtener información de %s: %v\n", path, err)
			return err
		}

		ui.Printf("\n📊 Información de %s:\n", path)
		ui.Printf("Nombre: %s\n", stat.Name())
		ui.Printf("Tamaño: %s\n", formatFileSize(stat.Size()))

		fileType := "Archivo"
		if stat.IsDir() {
			fileType = "Directorio"
		}
		ui.Printf("Tipo: %s\n", fileType)
		ui.Printf("Permisos: %s\n", stat.Mode().String())
		ui.Printf("Modificado: %s\n", stat.ModTime().Format("2006-01-02 15:04:05"))
		ui.Println("")
	}
	return nil
}
