package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
	"io"
	"os"
)

// GetCommand implements the get command
type GetCommand struct{}

func NewGetCommand() cli.ICommand {
	return &GetCommand{}
}

func (c *GetCommand) Name() string {
	return "get"
}

func (c *GetCommand) Help() string {
	return "Descarga un archivo del servidor remoto"
}

func (c *GetCommand) Usage() string {
	return "get <archivo_remoto> [archivo_local]"
}

func (c *GetCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: get <archivo_remoto> [archivo_local]")
		return fmt.Errorf("archivo remoto requerido")
	}

	client := session.GetSFTPClient()
	ui := session.GetUI()

	remotePath := session.GetCurrentDir() + "/" + args[0]
	localPath := "./local/" + args[0]
	if len(args) > 1 {
		localPath = args[1]
	}

	// Open remote file
	remoteFile, err := client.Open(remotePath)
	if err != nil {
		ui.Printf("❌ Error al abrir archivo remoto: %v\n", err)
		return err
	}

	defer remoteFile.Close()

	// Create local file
	localFile, err := os.Create(localPath)
	if err != nil {
		ui.Printf("❌ Error al crear archivo local: %v\n", err)
		return err
	}

	defer localFile.Close()

	// Copy content
	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		ui.Printf("❌ Error al descargar archivo: %v\n", err)
		return err
	}

	ui.Printf("✅ Archivo descargado: %s -> %s\n", remotePath, localPath)
	return nil
}
