package main

import (
	"erp/src/log"
	"erp/src/service"
	"time"
)

func main() {
	for {
		svc := service.NewService()
		err := svc.Start()
		if err != nil {
			log.Errorf("Failed to start service. error:[%s]",err.Error())
			time.Sleep(5*time.Second)
			continue
		}
		break
	}
	
	<-make(chan struct{})
}
