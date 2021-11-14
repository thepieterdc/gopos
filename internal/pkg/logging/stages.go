package logging

import "github.com/sirupsen/logrus"

func BootStage() logrus.Fields {
	return logrus.Fields{"stage": "boot"}
}

func RunningStage() logrus.Fields {
	return logrus.Fields{"stage": "running"}
}

func ShutdownStage() logrus.Fields {
	return logrus.Fields{"stage": "shutdown"}
}