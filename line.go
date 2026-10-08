package printer

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(0)
}

var verbose = false

func SetVerbose(value bool) {
	verbose = value
}

type Line interface {
	ExitProgram(value bool) Line
	WithIcon(icon string) Line
	WithFormat(color string) Line
	Debug(message string)
	PrintText(message string)
	PrintError(err error)
}

type line struct {
	exitProgram bool
	icon        string
	format      string
	args        []any
}

func (l *line) WithIcon(icon string) Line {
	l.icon = icon
	return l
}

func (l *line) WithFormat(value string) Line {
	l.format = value
	return l
}

func (l *line) ExitProgram(value bool) Line {
	l.exitProgram = value
	return l
}

func (l *line) PrintText(message string) {
	if len(l.icon) < 1 {
		l.format = "%s"
		l.args = []any{message}
	} else {
		l.format = "%s %s"
		l.args = []any{l.icon, message}
	}
	printLine(l)
}

func (l *line) PrintError(err error) {
	if len(l.icon) < 1 {
		l.format = "%v"
		l.args = []any{err}
	} else {
		l.format = "%s %v"
		l.args = []any{l.icon, err}
	}
	printLine(l)
}

func (l *line) Debug(message string) {
	if !verbose {
		return
	}

	l.PrintText(message)
}

func newLine() Line {
	return &line{
		exitProgram: false,
		format:      "",
		args:        []any{},
		icon:        "",
	}
}

func printLine(line *line) {
	format := line.format
	args := line.args
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")

	if line.exitProgram {
		log.Fatalf("%s "+format, append([]any{timestamp}, args...)...)
	}
	log.Printf("%s "+format, append([]any{timestamp}, args...)...)
}
