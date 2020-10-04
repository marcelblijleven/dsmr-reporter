package reporters

import (
	"github.com/marcelblijleven/dsmrreporter/dsmr"
	"log"
	"os"
)

type StdOutReporter struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewStdOutReporter() *StdOutReporter {
	reporter := &StdOutReporter{
		infoLog:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return reporter
}

func (s *StdOutReporter) Update(telegram dsmr.Telegram) {
	s.infoLog.Println(telegram)
}

func (s *StdOutReporter) Log(msg string) {
	s.infoLog.Println(msg)
}

func (s *StdOutReporter) Error(err error) {
	newErr := s.errorLog.Output(2, err.Error())

	if err != nil {
		s.errorLog.Println(newErr)
	}
}

func (s *StdOutReporter) Fatal(err error) {
	s.errorLog.Fatal(err)
}
