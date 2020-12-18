package tui

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var home string
var workDir string
var logFile string

func init() {
	// home
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("home dir not set: %s", err))
	}

	// work-dir
	workDir = home + "/.tui"
	if err := os.MkdirAll(workDir, 0777); err != nil {
		panic(fmt.Sprintf("mkdir work-dir: %s failed: %s", workDir, err))
	}

	// log
	logFile = workDir + "/log.log"
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		ForceColors:   true,
	})
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("init log file: %s failed: %s", logFile, err))
	}
	logrus.SetOutput(f)
	logrus.SetLevel(logrus.DebugLevel)
}
