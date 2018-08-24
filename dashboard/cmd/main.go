package main

import (
	"os"

	"github.com/plutov/surv/dashboard/pkg/api"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Out = os.Stdout

	srv := api.New(logger)
	if srv != nil {
		srv.Run()
	}
}
