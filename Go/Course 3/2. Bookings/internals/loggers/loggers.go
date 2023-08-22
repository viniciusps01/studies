package loggers

import (
	"log"
	"os"
)

type Loggers struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func New() Loggers {
	infoLogger := log.New(os.Stdin, "INFO\t", log.Ltime|log.Ldate)
	errorLogger := log.New(os.Stdin, "ERROR\t", log.Ltime|log.Ldate)

	l := Loggers{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}

	return l
}
