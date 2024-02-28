package service

import (
	"context"
	"grscan/api/models"
	"grscan/pkg/logger"
	"grscan/pkg/security"
	"grscan/storage"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewAuthService(storage storage.IStorage, log logger.ILogger) authService {
	return authService{
		storage: storage,
		log: log,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) error {
	password, err :=  a.storage.User().GetPasswordByLogin(ctx, loginRequest.Login)
	if err != nil {
		return err
	}

	if err = security.CompareHashAndPassword(password, loginRequest.Password); err != nil {
		return err
	}

	return nil
}