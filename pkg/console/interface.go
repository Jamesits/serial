package console

type InitFunc func() (in chan<- []byte, out <-chan []byte, err error)

var AvailableConsoleTypes = []InitFunc{
	ConsoleRawSyscall,
	ConsoleRawExperimental,
	ConsolePlain,
}
