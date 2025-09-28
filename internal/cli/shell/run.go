package shell

import (
	"bufio"
	"context"
	"fmt"
	"go-sftp-cli/internal/cli"
	"go-sftp-cli/kit/client/sftp"
	"go-sftp-cli/kit/client/ssh"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Interactive struct {
	registry cli.ICommandRegistry
	ui       cli.IUserInterface
}

func NewInteractive(registry cli.ICommandRegistry, ui cli.IUserInterface) *Interactive {
	return &Interactive{
		registry: registry,
		ui:       ui,
	}
}

func (in *Interactive) Run() {
	opts := in.promptConnectionOptions()

	sshClient, err := ssh.New(opts).Dial()
	if err != nil {
		log.Fatal(err)
		return
	}

	sftpClient, err := sftp.NewSFTPClient(sshClient)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		_ = sftpClient.Close()
	}()

	in.ui.Println("🚀 Conectado al servidor FTP/SFTP")

	in.next(sftpClient)
}

func (in *Interactive) next(sftpClient *sftp.Client) {

	// Initialize current directory
	currentDir, _ := sftpClient.GetWd()
	if currentDir == "" {
		currentDir = "/"
	}

	// Create session context
	session := cli.NewSessionContext(sftpClient, in.ui)
	session.SetCurrentDir(currentDir)

	in.ui.Println("Escribe 'help' para ver los comandos disponibles o 'exit' para salir")
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()

	for {
		in.ui.Printf("ftp:%s$ ", session.GetCurrentDir())

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		commandName := parts[0]
		args := parts[1:]

		// Handle special command
		if commandName == "exit" || commandName == "quit" {
			in.ui.Println("👋 Desconectando...")
			return
		}

		// Execute command through registry
		err := in.registry.Execute(ctx, commandName, session, args)
		if err != nil {
			in.ui.Printf("❌ Comando desconocido: %s\n", commandName)
			in.ui.Println("Escribe 'help' para ver los comandos disponibles")
		}
	}
}

func (in *Interactive) promptConnectionOptions() ssh.Options {
	scanner := bufio.NewScanner(os.Stdin)

	in.ui.Println("==========================================================================")
	in.ui.Println("🔐 Configuración de conexión FTP/SFTP")
	in.ui.Println("==========================================================================")

	// Solicitar Host
	in.ui.Print("Host: ")
	scanner.Scan()
	host := strings.TrimSpace(scanner.Text())

	// Solicitar Puerto
	var port int
	for {
		fmt.Print("Port (default 22): ")
		scanner.Scan()
		portStr := strings.TrimSpace(scanner.Text())
		if portStr == "" {
			port = 22
			break
		}
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil || port <= 0 || port > 65535 {
			in.ui.Println("❌ Puerto inválido. Debe ser un número entre 1 y 65535.")
			continue
		}
		break
	}

	// Solicitar Username
	in.ui.Print("Username: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	// Solicitar Password
	in.ui.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	in.ui.Printf("✅ Configuración completada para %s@%s:%d\n\n", username, host, port)

	return ssh.Options{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Timeout:  30 * time.Second,
		CiphersExtras: []string{
			"aes256-cbc",
			"aes128-cbc",
		},
		KeyExchangesExtras: []string{
			"diffie-hellman-group-exchange-sha1",
			"diffie-hellman-group-exchange-sha256",
		},
		MACsExtras: []string{},
	}
}
