package logger

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

var logfilePath string
var logFile *os.File

// Initialize the directory that will be used by the time tracker.
func Initialize(logfileName string) {
	directoryPath := initDirectory()
	logFile = InitLogfile(directoryPath, logfileName)
}

func initDirectory() string {

	path := filepath.Join(homeDirectory(), ".timewatch")
	os.MkdirAll(path, os.ModePerm)

	log.Println("Initialized directory ", path)

	return path
}

// InitLogfile initializes the logfile for the time tracker by either creating the file if it is absent or loading the existing file.
func InitLogfile(path string, filename string) *os.File {

	logfilePath = fullPath(path, filename)
	file, err := os.OpenFile(logfilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		logFile = newLogfile()
	}
	logFile = file
	return logFile
}

// Add an entry to the log file.
func Add(entry string) {

	defer logFile.Close()
	logFile.WriteString(entry)
	logFile.WriteString("\n")
}

func newLogfile() *os.File {

	file, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	file.WriteString("Date;Weekday;Begin;End;Duration\n")
	file.Sync()
	logFile = file
	return logFile
}

func homeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func fullPath(path string, filename string) string {

	match, _ := regexp.MatchString(".*/$", path)
	if !match {
		path = path + "/"
	}
	return path + filename
}
