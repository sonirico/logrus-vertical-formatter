package main

import (
	"github.com/sirupsen/logrus"
	logrusverticalformatter "github.com/sonirico/logrus-vertical-formatter"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(
		logrusverticalformatter.New(
			logrusverticalformatter.WithPadRight(" ", 40),
		),
	)

	logger := log.WithField("service", "orders")

	logger.Infoln("order placed successfully")
}
