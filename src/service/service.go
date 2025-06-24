package service

import (
	"erp/src/postgres"
	"erp/src/rest"
	"fmt"
)

type service struct {
}

func (s *service) Start() (err error){

	err = postgres.Init()
	if err != nil {
		return fmt.Errorf("failed to init postgres:[%w]",err)
	}

	err = rest.Init()
	if err != nil {
		return fmt.Errorf("failed to init rest server:[%w]",err)
	}

	return
}

func (s *service) Stop()error {
	return nil
}

func NewService() service{
	return service{}
}