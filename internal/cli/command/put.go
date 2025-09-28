package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
	"io"
	"os"
)

// PutCommand implements the put command
type PutCommand struct{}

func NewPutCommand() cli.ICommand {
	return &PutCommand{}
}

func (c *PutCommand) Name() string {
	return "put"
}

func (c *PutCommand) Help() string {
	return "Sube un archivo al servidor remoto"
}

func (c *PutCommand) Usage() string {
	return "put <archivo_local> [archivo_remoto]"
}

func (c *PutCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: put <archivo_local> [archivo_remoto]")
		return fmt.Errorf("archivo local requerido")
	}

	client := session.GetSFTPClient()
	ui := session.GetUI()

	localPath := args[0]
	remotePath := localPath
	if len(args) > 1 {
		remotePath = args[1]
	}

	// Open local file
	localFile, err := os.Open(localPath)
	if err != nil {
		ui.Printf("❌ Error al abrir archivo local: %v\n", err)
		return err
	}
	defer localFile.Close()

	// Create remote file
	remoteFile, err := client.Create(remotePath)
	if err != nil {
		ui.Printf("❌ Error al crear archivo remoto: %v\n", err)
		return err
	}
	defer remoteFile.Close()

	// Copy content
	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		ui.Printf("❌ Error al subir archivo: %v\n", err)
		return err
	}

	ui.Printf("✅ Archivo subido: %s -> %s\n", localPath, remotePath)
	return nil
}
