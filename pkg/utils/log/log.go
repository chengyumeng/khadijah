package log

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/fatih/color"
)

var (
	CmdLogger = log.New() // command line exec time log
	AppLogger = log.New() // application exec time log
)

func init() {
	CmdLogger.Formatter = &CmdFormatter{}
	CmdLogger.Out = os.Stdout
	AppLogger.Out = os.Stderr
}

type CmdFormatter struct{} // Command log formatter interface.

func (f *CmdFormatter) Format(entry *log.Entry) ([]byte, error) {
	var colorFunc func(string, ...interface{}) string
	switch entry.Level {
	case log.PanicLevel:
		colorFunc = color.BlueString
	case log.FatalLevel:
		colorFunc = color.MagentaString
	case log.ErrorLevel:
		colorFunc = color.RedString
	case log.WarnLevel:
		colorFunc = color.YellowString
	case log.InfoLevel:
		colorFunc = color.CyanString
	case log.DebugLevel:
		colorFunc = color.GreenString
	}
	if colorFunc != nil {
		return []byte(colorFunc("%s\n", entry.Message)), nil
	}
	return []byte(entry.Message + "\n"), nil
}

// Init a new command logger
func NewCmdLogger(module string) *log.Entry {
	return CmdLogger.WithField("module", module)
}

// Init a new application logger
func NewAppLogger(module string) *log.Entry {
	return AppLogger.WithField("module", module)
}
