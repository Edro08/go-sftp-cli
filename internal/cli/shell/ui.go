package shell

import (
	"bufio"
	"fmt"
	"go-sftp-cli/internal/cli"
	"os"
)

type consoleUI struct {
	scanner *bufio.Scanner
}

func NewConsoleUI() cli.IUserInterface {
	return &consoleUI{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *consoleUI) Print(message string) {
	fmt.Print(message)
}

func (c *consoleUI) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (c *consoleUI) Println(message string) {
	fmt.Println(message)
}

func (c *consoleUI) PrintError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
}

func (c *consoleUI) Prompt(message string) string {
	fmt.Print(message)
	if c.scanner.Scan() {
		return c.scanner.Text()
	}
	return ""
}

func (c *consoleUI) ReadLine() (string, error) {
	if c.scanner.Scan() {
		return c.scanner.Text(), nil
	}
	return "", c.scanner.Err()
}
