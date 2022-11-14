package main

import (
	"erp/postgres"
	"erp/rest"
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	zapLog, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	logger := zapLog.Sugar()
	logger.Info("Starting ERP system v3")

DB:
	err = postgres.Init()
	if err != nil {
		logger.Errorf("POSTGRES|INIT: %s", err.Error())
		time.Sleep(5 * time.Second)
		logger.Error("POSTGRES|INIT Retrying to connect...", err.Error())
		goto DB
	}

REST:
	err = rest.Init()
	if err != nil {
		logger.Errorf("REST|INIT: %s", err.Error())
		time.Sleep(5 * time.Second)
		logger.Error("REST|INIT Retrying to bring server up...", err.Error())
		goto REST
	}

	<-make(chan struct{})
}
