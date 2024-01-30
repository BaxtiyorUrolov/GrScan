package storage

import "grscan/api/models"

type IStorage interface {
	Close()
	User() IUserStorage
}

type IUserStorage interface {
	Create(models.CreateUser) (string, error)
	GetByID(models.PrimaryKey) (models.User, error)
	IsLoginExist(login string) (bool, error)
}


