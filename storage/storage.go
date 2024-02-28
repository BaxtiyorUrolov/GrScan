package storage

import (
	"context"
	"grscan/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	IsLoginExist(context.Context, string) (bool, error)
	GetPasswordByLogin(context.Context, string) (string, error)
}


