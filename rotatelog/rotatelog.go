package rotatelog
//package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RotateLog struct {
	logAbsPath  string
	logPreffix  string
	flag        int
	fp          *os.File
	logger      *log.Logger
	rotateTime  time.Duration
	rotateMutex sync.Mutex // protect rotate file
	//rotateMutex sync.RWMutex
}

// NewRotateLog creat a object of RotateLog
func NewRotateLog(path string, prefix string, flag int) (*RotateLog, error) {
	if len(path) == 0 {
		return nil, errors.New("NewRotateLog path is nil")
	}
	if len(prefix) != 0 {
		prefix += " "
	}
	absPath, err := filepath.Abs(path)
	var fp *os.File
	fp, err = os.OpenFile(absPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("NewRotateLog open abs path err: ", err.Error())
	}
	return &RotateLog{
		logAbsPath: absPath,
		logPreffix: prefix,
		fp:         fp,
		flag:       flag,
		logger:     log.New(fp, prefix, flag),
		rotateTime: time.Duration(time.Hour * 24),
	}, nil
}

// Println RotateLog println info
func (r *RotateLog) Println(args ...interface{}) {
	r.rotateMutex.Lock()
	defer r.rotateMutex.Unlock()

	r.logger.Println(args)
}

func (r *RotateLog) SetRotateTime(t time.Duration) {
	r.rotateTime = t
}

func (r *RotateLog) RotateForce() {
	if stat, _ := os.Stat(r.logAbsPath); stat.Size() == 0 {
		// empty file don't rotate
		return
	}

	suffix := time.Now().Format("2006-01-02")	// default rotate with per day
	rotateFile := r.logAbsPath + "." + suffix

	_, err := os.Stat(rotateFile)
	if os.IsNotExist(err) {
		if err = os.Rename(r.logAbsPath, rotateFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	r.fp.Close()

	r.fp, _ = os.OpenFile(r.logAbsPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	r.logger.SetOutput(r.fp)
}

func (r *RotateLog) rotate() {
	r.rotateMutex.Lock()
	defer r.rotateMutex.Unlock()

	r.RotateForce()
}

func (r *RotateLog) RotateWithTime() {
	go func() {
		for {
			time.Sleep(r.rotateTime)
			r.rotate()
		}
	}()
}

func xmain() {
	rl, err := NewRotateLog("logs/rotatelog", "[NewRotate]", log.LUTC|log.LstdFlags)
	if err != nil {
		fmt.Println("get new rotate log err: ", err)
	}
	rl.SetRotateTime(time.Duration(time.Second * 5))
	rl.RotateWithTime()
	for i:=0;i<20;i++{
		time.Sleep(time.Second)
		rl.logger.Println("12345")
		rl.logger.Println("12345")
		rl.logger.Println("12345")
		rl.logger.Println("12345")
		rl.logger.Println("12345")
	}
}
