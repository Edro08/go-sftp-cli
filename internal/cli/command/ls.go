package command

import (
	"context"
	"go-sftp-cli/internal/cli"
	"path/filepath"
)

// LsCommand implements the ls command
type LsCommand struct{}

func NewLsCommand() cli.ICommand {
	return &LsCommand{}
}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Help() string {
	return "Lista el contenido del directorio"
}

func (c *LsCommand) Usage() string {
	return "ls [directorio]"
}

func (c *LsCommand) Execute(ctx context.Context, session *cli.SessionContext, args []string) error {
	client := session.GetSFTPClient()
	ui := session.GetUI()
	currentDir := session.GetCurrentDir()

	path := currentDir
	if len(args) > 0 {
		if filepath.IsAbs(args[0]) {
			path = args[0]
		} else {
			path = filepath.Join(currentDir, args[0])
		}
	}

	files, err := client.ReadDir(path)
	if err != nil {
		ui.Printf("❌ Error al listar directorio: %v\n", err)
		return err
	}

	ui.Printf("\n📁 Contenido de %s:\n", path)
	for _, file := range files {
		var fileType string
		var size string

		if file.IsDir() {
			fileType = "📁"
			size = "<DIR>"
		} else {
			fileType = "📄"
			size = formatFileSize(file.Size())
		}

		ui.Printf("%s %-20s %10s %s\n", fileType, file.Name(), size, file.ModTime().Format("2006-01-02 15:04:05"))
	}
	ui.Println("")
	return nil
}
