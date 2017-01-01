package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"
)

type RotateLog struct {
	logPath    string
	logPreffix string
	flag       int
	fp         *os.File
	logger     *log.Logger
	rorateTime time.Duration
}

// NewRotateLog creat a object of RotateLog
func NewRotateLog(path string, prefix string, flag int) (*RotateLog, error) {
	if len(path) == 0 {
		return nil, errors.New("NewRotateLog path is nil")
	}
	absPath, err := filepath.Abs(path)
	var fp *os.File
	fp, err = os.OpenFile(absPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("NewRotateLog open abs path err: ", err.Error())
	}
	return &RotateLog{
		logPath:    path,
		logPreffix: prefix,
		fp:         fp,
		logger:     log.New(fp, prefix, flag),
		rorateTime: time.Duration(time.Hour * 24),
	}, nil
}

// Println RotateLog println info
func (r *RotateLog) Println(args ...interface{}) {
	r.logger.Println(args)
}

func main() {
	log.Println("aabbcc")
}
