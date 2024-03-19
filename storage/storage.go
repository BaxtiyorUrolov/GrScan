package storage

import (
	"context"
	"grscan/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Register() IRegisterStorage
}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	IsLoginExist(context.Context, string) (bool, error)
	GetPasswordByLogin(context.Context, string) (string, error)
}


type IRegisterStorage interface {
	Create(context.Context, models.CreateRegister) error
	GetByID(context.Context, string) (models.Register, error)
	UpdateStatus(context.Context, string) error
}