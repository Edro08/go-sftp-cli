package cli

type IShell interface {
	Run()
}

type IUserInterface interface {
	Print(message string)
	Printf(format string, args ...interface{})
	Println(message string)
	PrintError(err error)
	Prompt(message string) string
	ReadLine() (string, error)
}
