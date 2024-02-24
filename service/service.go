package service

import (
	"grscan/pkg/logger"
	"grscan/storage"
)

type IServiceManager interface {
	User() userService
}

type Service struct {
	userService userService
}

func New(storage storage.IStorage,  log logger.ILogger) Service {
	services := Service{}

	services.userService = NewUserService(storage, log)

	return  services
}

func (s Service) User() userService {
	return s.userService
}
