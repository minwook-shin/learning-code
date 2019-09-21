package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	logrus.SetReportCaller(false)

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.TraceLevel)

	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Trace("Trace")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Debug("Debug")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Info("Info")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Warn("Warn")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Error("Error")

	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Fatal("Fatal")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Panic("Panic")

	entry := logrus.WithFields(logrus.Fields{"Fields": "TEST"})
	entry.Trace("Trace logger")
	entry.Debug("Debug logger")
	entry.Info("Info logger")
	entry.Warn("Warn logger")
	entry.Error("Error logger")

	entry.Fatal("Fatal logger")
	entry.Panic("Panic logger")

}
