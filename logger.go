package easylog

import (
    "fmt"
    "io"
    "os"
    "strings"
    "sync"
    "time"
)

const (
    DEBUG = iota
    INFO
    WARN
    ERROR
)

var levelName = map[uint]string{
    DEBUG: "DEBUG",
    INFO:  "INFO",
    WARN:  "WARN",
    ERROR: "ERROR",
}

type Logger struct {
    FilePath string
    Writer   io.Writer
    Level    uint
    Mutex    sync.Mutex
}

func NewLogger(fp string, level uint) *Logger {
    // create dir,
    if strings.ContainsAny(fp, "/\\") {
        lastIdx := strings.LastIndexAny(fp, "/\\")
        dirPath := fp[:lastIdx]

        if _, err := os.Stat(dirPath); os.IsNotExist(err) {
            err := os.MkdirAll(dirPath, 0700)
            if err != nil {
                panic("Error while create log dir: " + err.Error())
            }
        }
    }

    // open log file,
    file, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        panic("Error while open log file: " + err.Error())
    }

    return &Logger{
        FilePath: fp,
        Writer:   io.MultiWriter(os.Stdout, file),
        Level:    level,
        Mutex:    sync.Mutex{},
    }
}

func (logger *Logger) Log(level uint, format string, a ...interface{}) {
    if level < logger.Level {
        return
    }

    logger.Mutex.Lock()
    preFormat := "%s [%s] "
    preA := []interface{}{time.Now().Format(time.RFC3339), levelName[level]}

    _, err := fmt.Fprintf(logger.Writer, preFormat+format+"\n", append(preA, a...)...)
    if err != nil {
        panic("Failed to log: " + err.Error())
    }
    logger.Mutex.Unlock()
}

func (logger *Logger) Debug(format string, a ...interface{}) {
    logger.Log(DEBUG, format, a...)
}

func (logger *Logger) Info(format string, a ...interface{}) {
    logger.Log(INFO, format, a...)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
    logger.Log(WARN, format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
    logger.Log(ERROR, format, a...)
}
