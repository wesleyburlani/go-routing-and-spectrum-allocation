package logs

import (
	"log"
	"os"
)

type FileLogger struct {
	File *os.File
}

func NewFileLogger(filePath string) *FileLogger {
	os.Remove(filePath)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	return &FileLogger{
		File: f,
	}
}

func (f *FileLogger) Log(logs ...interface{}) {
	log.New(f.File, "", log.LstdFlags|log.Lshortfile).Print(logs...)
}
