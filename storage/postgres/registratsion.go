package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"grscan/api/models"
	"grscan/pkg/logger"
	"grscan/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type registerRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewRegisterRepo(db *pgxpool.Pool, log logger.ILogger) storage.IRegisterStorage {
	return &registerRepo{
		db:  db,
		log: log,
	}
}

func (r *registerRepo) Create(ctx context.Context, createUser models.CreateRegister) error {
	uid := uuid.New()
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
		INSERT INTO verify_cods (id, phone, code, created_at) 
		VALUES ($1, $2, $3, $4)
		`,
		uid,
		createUser.Phone,
		createUser.Code,
		createdAt,
	)
	if err != nil {
		r.log.Error("error is while inserting code and phone", logger.Error(err))
	}

	return nil
}

func (r *registerRepo) GetByID(ctx context.Context, phone string) (models.Register, error) {
    user := models.Register{}
	var createdAt sql.NullString

    fmt.Println("Phone: ", phone)

    query := `
        SELECT code, created_at FROM verify_cods WHERE phone = $1`
    err := r.db.QueryRow(ctx, query, phone).Scan(
        &user.Code,
        &createdAt,
    )
    if err != nil {
        r.log.Error("error while selecting code by phone", logger.Error(err))
        return user, err
    }

	if createdAt.Valid {
		user.CreatedAT = createdAt.String
	}


    return user, nil
}


func (r *registerRepo) UpdateStatus(ctx context.Context, phone string) error {

	query := `update users set user_verify = $1 where phone = $2`

	if _, err := r.db.Exec(ctx, query, true, phone); err != nil {
		r.log.Error("error is while updeting user verify status", logger.Error(err))
		return err
	}

	return nil
}
