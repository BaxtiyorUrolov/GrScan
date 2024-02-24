package service

import (
	"context"
	"grscan/api/models"
	"grscan/pkg/check"
	"grscan/pkg/logger"
	"grscan/pkg/security"
	"grscan/pkg/sms"
	"grscan/storage"
)

type userService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewUserService(storage storage.IStorage, log  logger.ILogger) userService {
	return userService{
		storage: storage,
		log: log,
	}
}

func (u userService) Create(ctx context.Context, createUser models.CreateUser) (models.User, error) {
	u.log.Info("user create service layer", logger.Any("user", createUser))

	var err error

	if !check.PhoneNumber(createUser.Phone) {
		u.log.Error("Incorrect phone number", logger.Error(err))
		return models.User{}, nil
	}

	if !check.ValidatePassword(createUser.Password) {
		u.log.Error("Invalid password", logger.Error(err))
		return models.User{}, nil
	}

	if exists, err := check.IsLoginExist(createUser.Login, u.storage.User()); err != nil {
		u.log.Error("Error while checking login existence", logger.Error(err))
		return models.User{}, err
	} else if exists {
		u.log.Error("Login already exists", logger.Error(err)) 
		return models.User{}, err
	}

	password, err := security.HashPassword(createUser.Password)
	if err != nil {
		u.log.Error("Error while hashing password", logger.Error(err))
		return models.User{}, err
	}
	createUser.Password = password

	code := sms.GenerateVerificationCode()

	if err := sms.Send(createUser.Phone, code); err != nil {
		u.log.Error("Error while sending verification code", logger.Error(err))
		return models.User{}, err
	}

	pKey, err := u.storage.User().Create(context.Background(), createUser)
	if err != nil {
		u.log.Error("Error while creating user", logger.Error(err))
		return models.User{}, err
	}

	user, err := u.storage.User().GetByID(context.Background(), models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		u.log.Error("Error in service layer when getting user by id", logger.Error(err))
		return user, err
	}

	return user, nil
}