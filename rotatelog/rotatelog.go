package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"
	"fmt"
	"sync"
)

type RotateLog struct {
	logAbsPath     string
	logPreffix string
	flag       int
	fp         *os.File
	logger     *log.Logger
	rorateTime time.Duration
	rotateMutex sync.RWMutex
}

// NewRotateLog creat a object of RotateLog
func NewRotateLog(path string, prefix string, flag int) (*RotateLog, error) {
	if len(path) == 0 {
		return nil, errors.New("NewRotateLog path is nil")
	}
	if len(prefix) != nil {
		prefix += " "
	}
	absPath, err := filepath.Abs(path)
	var fp *os.File
	fp, err = os.OpenFile(absPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("NewRotateLog open abs path err: ", err.Error())
	}
	return &RotateLog{
		logAbsPath:    absPath,
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

func (r *RotateLog) SetRotateTime(t time.Duration) {
	r.rorateTime = t
}

func (r *RotateLog) rotateWithLock() {
	r.rotateMutex.Lock()
	defer r.rotateMutex.Unlock()

	if stat. err := os.Stat(r.logAbsPath)
	// empty file don't rotate
	if stat.Size() == 0 {
		return
	}
}

func main() {
	rl, err := NewRotateLog("logs/rotatelog", "[NewRotate]", log.LUTC|log.LstdFlags)
	if err != nil {
		fmt.Println("get new rotate log err: ", err)
	}
	rl.logger.Println("12345")
}
