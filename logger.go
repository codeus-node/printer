package printer

import di "github.com/codeus-node/generic-di"

func init() {
	di.Injectable(newLogger)
}

type Logger interface {
	Log() Line
	LogAndExit() Line
}

type logger struct {
}

func (l *logger) Log() Line {
	return newLine().
		ExitProgram(false)
}

func (l *logger) LogAndExit() Line {
	return newLine().
		ExitProgram(true)
}

func newLogger() Logger {
	return &logger{}
}
