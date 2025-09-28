package command

import (
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
	"path"
)

// CdCommand implements the cd command
type CdCommand struct{}

func NewCdCommand() cli.ICommand {
	return &CdCommand{}
}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) Help() string {
	return "Cambia el directorio actual"
}

func (c *CdCommand) Usage() string {
	return "cd <directorio>"
}

func (c *CdCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	if len(args) == 0 {
		session.GetUI().Println("❌ Uso: cd <directorio>")
		return fmt.Errorf("directorio requerido")
	}

	dirArg := args[0] // Porque args[0] es "cd"

	client := session.GetSFTPClient()
	ui := session.GetUI()
	currentDir := session.GetCurrentDir()

	var newPath string
	if path.IsAbs(dirArg) {
		newPath = dirArg
	} else {
		newPath = path.Join(currentDir, dirArg)
	}

	// Check if directory exists
	stat, err := client.Stat(newPath)
	if err != nil {
		ui.Printf("❌ Error al acceder al directorio: %v\n", err)
		return err
	}

	if !stat.IsDir() {
		ui.Printf("❌ %s no es un directorio\n", newPath)
		return fmt.Errorf("%s no es un directorio", newPath)
	}

	session.SetCurrentDir(newPath)
	return nil
}

