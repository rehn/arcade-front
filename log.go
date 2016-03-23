package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	logfile *os.File
)

func LogSetup() {
	var err error
	fil := path.Base(os.Args[0])
	dt := time.Now().Format("2006-01-02")
	p, err := PathToABS("log/" + fil + "-" + dt + ".log")
	LogFatal(err)
	logfile, err = os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	LogFatal(err)
	log.SetOutput(logfile)
	go LogCleaner()
}

func LogCleaner() {
	p, err := PathToABS("log/")
	LogFatal(err)
	for {
		files, err := filepath.Glob(p + "*.log")
		if LogError(err) {
			continue
		}
		for _, file := range files {
			info, err := os.Stat(file)
			if err != nil {
				continue
			}
			d := time.Now().Sub(info.ModTime())
			if d > (48 * time.Hour) {
				LogError(os.Remove(file))
			}
		}
		time.Sleep(20 * time.Hour)
	}
}

func LogDefer() {
	logfile.Close()
}

func LogMsg(msg string) {
	log.SetPrefix("[MESSAGE]")
	log.Print(msg)
}

func LogInfo(err error) {
	if err != nil {
		log.SetPrefix("[INFO]")
		log.Print(err)
	}
}
func LogError(err error) bool {
	if err != nil {
		log.SetPrefix("[ERROR]")
		log.Print(err)
		return true
	}
	return false
}

func LogWarning(err error) bool {
	if err != nil {
		log.SetPrefix("[WARNING]")
		log.Print(err)
		return true
	}
	return false
}

func LogFatal(err error) bool {
	if err != nil {
		log.SetPrefix("[Fatal]")
		log.Fatal(err)
		return true
	}
	return false
}
