package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type RotateFile struct {
	path string
	file *os.File
	cron *cron.Cron
	mu   sync.Mutex
}

func NewRotateFile(path string) *RotateFile {
	rf := &RotateFile{
		path: filepath.Clean(path),
		cron: cron.New(),
	}
	rf.rotate()

	rf.cron.AddFunc("0 0 * * *", func() {
		rf.rotate()
	})
	rf.cron.Start()
	return rf
}

func (rf *RotateFile) Write(p []byte) (n int, err error) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.file.Write(p)
}

func (rf *RotateFile) rotate() {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	if _, err := os.Stat(rf.path); os.IsNotExist(err) {
		if err := os.MkdirAll(rf.path, 0755); err != nil {
			panic(fmt.Sprintf("Failed to create log directory: %s with error %v", rf.path, err))
		}
	}

	if rf.file != nil {
		if err := rf.file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
			rf.file = nil
		}
	}

	ymd := time.Now().Format("20060102")
	filename := fmt.Sprintf("%s/%s.log", rf.path, ymd)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %s with error %v", filename, err))
	}

	rf.file = file
}

func (rf *RotateFile) Close() {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	if rf.file != nil {
		rf.file.Close()
	}
	rf.cron.Stop()
}
