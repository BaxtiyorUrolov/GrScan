package postgres

import (
	"context"
	"fmt"
	"grscan/api/models"
	"grscan/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.IUserStorage {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(createUser models.CreateUser) (string, error) {
	uid := uuid.New()
	createdAt := time.Now()

	_, err := u.db.Exec(context.Background(), `
		INSERT INTO users (id, phone, login, password, user_type, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
		`,
		uid,
		createUser.Phone,
		createUser.Login,
		createUser.Password,
		createUser.UserType,
		createdAt,
	)
	if err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (u *userRepo) GetByID(pKey models.PrimaryKey) (models.User, error) {
	user := models.User{}

	query := `
		select user_id, phone, user, balance from users where id = $1 and user_role = 'customer'
`
	if err := u.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&user.UserID,
		&user.Phone,
		&user.Login,
		&user.Balance,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.User{}, err
	}

	return user, nil
}
