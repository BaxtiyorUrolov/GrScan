package service

import (
	"grscan/pkg/logger"
	"grscan/storage"
)

type IServiceManager interface {
	User() userService
	Auth() authService
	Register() registerService
}

type Service struct {
	userService userService
	authService authService
	registerService registerService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}

	services.userService = NewUserService(storage, log)
	services.authService = NewAuthService(storage, log)
	services.registerService = NewRegisterService(storage, log)

	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Auth() authService {
	return s.authService
}

func (s Service) Register() registerService{
	return s.registerService
}
