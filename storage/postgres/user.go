package postgres

import (
	"context"
	"fmt"
	"grscan/api/models"
	"grscan/pkg/logger"
	"grscan/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
	log logger.ILogger
}

func NewUserRepo(db *pgxpool.Pool, log logger.ILogger) storage.IUserStorage {
	return &UserRepo{
		db: db,
		log: log,
	}
}

func (u *UserRepo) Create(ctx context.Context, createUser models.CreateUser) (string, error) {
	uid := uuid.New()
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := u.db.Exec(ctx, `
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
		u.log.Error("error is while inserting data", logger.Error(err))
	}

	return uid.String(), nil
}

func (u *UserRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	user := models.User{}

	query := `
		SELECT id, user_id, phone, login, balance FROM users WHERE id = $1 AND user_type = 'customer'
	`
	if err := u.db.QueryRow(ctx, query, pKey.ID).Scan(
		&user.ID,
		&user.UserID,
		&user.Phone,
		&user.Login,
		&user.Balance,
	); err != nil {
		u.log.Error("error is while selecting user by id", logger.Error(err))
	}

	return user, nil
}

func (u *UserRepo) IsLoginExist(login string) (bool, error) {
	var exists bool
	err := u.db.QueryRow(context.Background(), `
		SELECT EXISTS (SELECT 1 FROM users WHERE login = $1)
	`, login).Scan(&exists)
	if err != nil {
		fmt.Println("error while checking login existence:", err)
		return false, err
	}

	return exists, nil
}