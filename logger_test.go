package easylog

import (
    "os"
    "testing"
)

func TestLogger(t *testing.T) {
    filepath := "easy.log"
    logger := NewLogger(filepath, INFO)

    logger.Debug("hello, %s", "world")
    logger.Info("hello, %s", "world")
    logger.Warn("hello, %s", "world")
    logger.Error("hello, %s", "world")
}

func TestLogger_FileRemoved(t *testing.T) {
    filepath := "easy_removed.log"
    logger := NewLogger(filepath, INFO)
    os.Remove(filepath)

    logger.Info("hello, %s", "world")
    logger.Warn("hello, %s", "world")
}

func TestLogger_InvalidArgs(t *testing.T) {
    filepath := "easy_invalid_args.log"
    logger := NewLogger(filepath, INFO)

    logger.Info("hello, %s")
    logger.Warn("hello, %s", 100)
}

func TestLogger_LogDir_relative(t *testing.T) {
    filepath := "log/easy.log"
    logger := NewLogger(filepath, INFO)

    logger.Info("hello, %s", "world")
}
func TestLogger_LogDir_abstract(t *testing.T) {
    filepath := "/tmp/log/easy.log"
    logger := NewLogger(filepath, INFO)

    logger.Info("hello, %s", "world")
}

func TestLogger_LogDir_windows(t *testing.T) {
    filepath := "D:\\log\\easy.log"
    logger := NewLogger(filepath, INFO)

    logger.Info("hello, %s", "world")
}
